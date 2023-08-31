package v1

import (
	"avito-backend-task/internal/repo/pgdb"
	"github.com/gofiber/fiber/v2"
)

func newSegmentRoutes(router *fiber.Router, r *pgdb.SegmentRepo) {
	(*router).Post("/create_segment", r.CreateSegment)
	(*router).Delete("/delete_segment/:id", r.DeleteSegment)
}
