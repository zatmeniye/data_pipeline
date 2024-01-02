package http

import (
	"bi/internal/dto"
	"bi/internal/service"
	"bi/pkg/keeper"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type sourceRouter struct {
	sourceService *service.SourceService
}

func withSourceRouter(app *fiber.App, sourceService *service.SourceService) {
	r := sourceRouter{sourceService: sourceService}
	source := app.Group("/source/")
	source.Get("/", r.getAll)
	source.Post("/", r.add)
	source.Post("/exec/", r.executeQuery)
}

// @tags	источник
// @param	source	body	dto.SourceAddDto	true	"source"
// @router	/source/ [post]
func (r sourceRouter) add(ctx *fiber.Ctx) error {
	var addDto dto.SourceAddDto

	err := ctx.BodyParser(&addDto)
	if err != nil {
		return err
	}

	var sourceId uint32

	if sourceId, err = r.sourceService.Add(ctx.Context(), addDto); err != nil {
		return err
	}

	return ctx.SendString(fmt.Sprintf("%d", sourceId))
}

// @tags	источник
// @param	query	body	dto.QueryDto	true	"query"
// @router	/source/exec/ [post]
func (r sourceRouter) executeQuery(ctx *fiber.Ctx) error {
	var queryDto dto.QueryDto

	err := ctx.BodyParser(&queryDto)
	if err != nil {
		return err
	}

	var dataset keeper.Dataset

	if dataset, err = r.sourceService.ExecuteQuery(ctx.Context(), queryDto); err != nil {
		return err
	}

	return ctx.JSON(dataset)
}

// @tags		источник
// @success	200	{array}	dto.SourceDto
// @router		/source/ [get]
func (r sourceRouter) getAll(ctx *fiber.Ctx) error {
	sources, err := r.sourceService.GetAll(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(sources)
}
