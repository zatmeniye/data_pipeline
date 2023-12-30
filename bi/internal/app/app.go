package app

import (
	"bi/config"
	"bi/internal/controller/http"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Run(l *zap.SugaredLogger, cfg *config.Config) error {
	app := fiber.New()

	http.WithRouter(l, app)

	return app.Listen(cfg.Http.Address)
}
