package main

import (
	"bi/config"
	"bi/internal/app"
	"bi/pkg/logger"
)

func main() {
	l := logger.New()

	cfg, err := config.New()
	if err != nil {
		l.Fatalln(err)
	}

	l.Fatalln(app.Run(l, cfg))
}
