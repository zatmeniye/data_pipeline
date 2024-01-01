package http

import (
	_ "bi/docs"
	"bi/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func WithRouter(
	app *fiber.App,
	sourceTypService *service.SourceTypService,
	sourceService *service.SourceService,
) {
	app.Get("/swagger/*", swagger.HandlerDefault)
	withSourceTypRouter(app, sourceTypService)
	withSourceRouter(app, sourceService)
}
