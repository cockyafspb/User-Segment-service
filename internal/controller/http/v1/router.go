package v1

import (
	"avito-backend-task/internal/repo/pgdb"
	"github.com/gofiber/fiber/v2"
)

type Repos struct {
	SegmentRepo *pgdb.SegmentRepo
	UserRepo    *pgdb.UserRepo
}

func NewRouter(app *fiber.App, repos Repos) {
	api := app.Group("/api").Group("v1")

	newSegmentRoutes(&api, repos.SegmentRepo)
	newUserRoutes(&api, repos.UserRepo)
}
