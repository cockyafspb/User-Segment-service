package pgdb

import (
	entity2 "avito-backend-task/internal/entity"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) CreateUser(context *fiber.Ctx) error {
	user := entity2.User{}
	err := context.BodyParser(&user)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	err = r.DB.Create(&user).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create a user"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user has been added",
		"data":    user,
	})
	return nil
}

func (r *UserRepo) DeleteUser(context *fiber.Ctx) error {
	user := entity2.User{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(user, id).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete a user",
		})
		return err
	}

	pair := entity2.UserSegmentPair{}
	err = r.DB.Where("user_id = ?", id).Delete(&pair).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete an info about user's segments",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user deleted successfully",
	})
	return nil
}

// TODO: потенциальная ошибка в связи с тем, что без id будут новый создаваться
func (r *UserRepo) AddSegments(context *fiber.Ctx) error {
	segments := entity2.SegmentsList{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}
	err := context.BodyParser(&segments)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "incorrect id"})
	}

	for _, v := range segments.Segments {
		s := entity2.Segment{}
		err = r.DB.Table("segments").Where("id = ?", v.ID).Where("slug = ?", v.Slug).Find(&s).Error
		if err != nil {
			context.Status(http.StatusBadRequest).JSON(
				&fiber.Map{"message": "no such segment here"})
			return err
		}
		pair := entity2.UserSegmentPair{UserID: uint(userId), SegmentID: v.ID}
		err := r.DB.Create(&pair).Error
		if err != nil {
			context.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"message": "could not add segments to user",
			})
			return err
		}
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "segments has been added to user",
		"user":    userId,
		"data":    segments.Segments,
	})

	return nil
}

// TODO: потенциальная ошибка в связи с тем, что без id будут новый создаваться
func (r *UserRepo) RemoveSegments(context *fiber.Ctx) error {
	segments := entity2.SegmentsList{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}
	err := context.BodyParser(&segments)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "incorrect id"})
	}
	for _, v := range segments.Segments {
		err = r.DB.Where("id = ? and slug = ?", v.ID, v.Slug).Find(&entity2.Segment{}).Error
		if err != nil {
			context.Status(http.StatusBadRequest).JSON(
				&fiber.Map{"message": "no such segment here"})
		}
		pair := entity2.UserSegmentPair{}
		err := r.DB.Where("user_id = ? and segment_id = ?", uint(userId), v.ID).Delete(&pair).Error
		if err != nil {
			context.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"message": "could not remove segments from user",
			})
			return err
		}
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "segments has been removed from user",
		"user":    userId,
		"data":    segments.Segments,
	})
	return nil
}

func (r *UserRepo) GetSegments(context *fiber.Ctx) error {
	id := context.Params("id")
	segments := &[]entity2.UserSegmentPair{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Where("user_id = ?", id).Find(&segments).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get a segments of user"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "segments of user " + id + " fetched successfully",
		"data":    segments,
	})

	return nil
}
