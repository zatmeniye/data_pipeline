package http

import (
	"bi/internal/service"
	"github.com/gofiber/fiber/v2"
)

type sourceTypRouter struct {
	sourceTypService *service.SourceTypService
}

func withSourceTypRouter(
	app *fiber.App,
	sourceTypService *service.SourceTypService,
) {
	r := sourceTypRouter{sourceTypService: sourceTypService}
	sourceTyp := app.Group("/source_typ/")
	sourceTyp.Get("/", r.getAll)
}

// @tags тип источника
// @success 200 {array} dto.SourceTypDto
// @router /source_typ/ [get]
func (r sourceTypRouter) getAll(ctx *fiber.Ctx) error {
	sourceTypes, err := r.sourceTypService.GetAll(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(sourceTypes)
}
