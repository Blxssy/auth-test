package main

import (
	"github.com/Blxssy/auth-test/internal/logger"
	"github.com/Blxssy/auth-test/internal/router"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"os"

	"github.com/Blxssy/auth-test/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}

	cfg := config.InitConfig()

	l := logger.SetupLogger(cfg.Env)
	l.Info("Logger", slog.String("Init logger", "successfully init logger"))

	r := gin.Default()
	router.InitRouter(r)

	// init db

	r.Run(cfg.Address)
}
