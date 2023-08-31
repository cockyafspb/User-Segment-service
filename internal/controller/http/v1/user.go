package v1

import (
	"avito-backend-task/internal/repo/pgdb"
	"github.com/gofiber/fiber/v2"
)

func newUserRoutes(router *fiber.Router, r *pgdb.UserRepo) {
	(*router).Post("/create_user", r.CreateUser)
	(*router).Delete("/delete_user/:id", r.DeleteUser)
	(*router).Get("/get_segments/:id", r.GetSegments)
	(*router).Post("/add_segments/:id", r.AddSegments)
	(*router).Delete("/remove_segments/:id", r.RemoveSegments)
}
