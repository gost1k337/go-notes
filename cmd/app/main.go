package main

import (
	"log"
	"tg_bot/configs"
	"tg_bot/internal/app"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
