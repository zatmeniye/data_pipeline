package app

import (
	"bi/config"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Run(l *zap.SugaredLogger, cfg *config.Config) error {
	app := fiber.New()

	return app.Listen(cfg.Http.Address)
}
