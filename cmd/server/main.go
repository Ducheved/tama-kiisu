package server

import (
	"github.com/Ducheved/tama-kiisu/internal/app"
	"github.com/Ducheved/tama-kiisu/internal/config"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	app, err := app.New(cfg)
	if err != nil {
		return err
	}

	return app.Run()
}
