package pgdb

import (
	entity2 "avito-backend-task/internal/entity"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type SegmentRepo struct {
	DB *gorm.DB
}

func (r *SegmentRepo) CreateSegment(context *fiber.Ctx) error {
	segment := entity2.Segment{}
	err := context.BodyParser(&segment)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	err = r.DB.Create(&segment).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create a segment"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "segment has been added",
		"data":    segment})
	return nil
}

func (r *SegmentRepo) DeleteSegment(context *fiber.Ctx) error {
	segment := entity2.Segment{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(segment, id).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete a segment",
		})
		return err
	}

	pair := &[]entity2.UserSegmentPair{}
	err = r.DB.Where("segment_id = ?", id).Delete(&pair).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete an info about users of this segment",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "segment deleted successfully",
	})
	return nil
}
