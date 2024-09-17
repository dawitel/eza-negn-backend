package main

import (
	"log"

	"github.com/dawitel/eza-negn-backend/pkg/config"
)

const (
	env = "dev"
)

func main() {
	configManager := config.NewConfigManager(env)
	cfg := configManager.GetConfig()
	log.Printf("Starting server on %s:%s in %s mode", cfg.Server.Host, cfg.Server.Port, cfg.Server.Mode)
}