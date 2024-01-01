package app

import (
	"bi/config"
	"bi/internal/controller/http"
	"bi/internal/mapper"
	"bi/internal/repository"
	"bi/internal/service"
	"bi/pkg/database"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Run(l *zap.SugaredLogger, cfg *config.Config) error {
	app := fiber.New()

	embeddedDb, err := database.New(cfg.Database)
	if err != nil {
		return err
	}

	var (
		sourceTypRepository = repository.NewSourceTypRepository(embeddedDb)
		sourceRepository    = repository.NewSourceRepository(embeddedDb)
	)

	var (
		sourceTypMapper = mapper.NewSourceTypMapper()
		sourceMapper    = mapper.NewSourceMapper(sourceTypMapper, sourceTypRepository)
	)

	var (
		sourceTypService = service.NewSourceTypService(l, sourceTypRepository, sourceTypMapper)
		sourceService    = service.NewSourceService(l, sourceRepository, sourceMapper)
	)

	http.WithRouter(
		app,
		sourceTypService,
		sourceService,
	)

	return app.Listen(cfg.Http.Address)
}
