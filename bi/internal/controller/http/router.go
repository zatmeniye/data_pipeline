package http

import (
	"bi/internal/service"
	"github.com/gofiber/fiber/v2"
)

func WithRouter(
	app *fiber.App,
	sourceTypService *service.SourceTypService,
) {
	withSourceTypRouter(app, sourceTypService)
}
