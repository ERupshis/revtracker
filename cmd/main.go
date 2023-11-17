package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/erupshis/revtracker/internal/auth"
	"github.com/erupshis/revtracker/internal/auth/data"
	"github.com/erupshis/revtracker/internal/auth/jwtgenerator"
	usersStorage "github.com/erupshis/revtracker/internal/auth/users/storage"
	"github.com/erupshis/revtracker/internal/config"
	"github.com/erupshis/revtracker/internal/controller"
	"github.com/erupshis/revtracker/internal/db"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	reformManager "github.com/erupshis/revtracker/internal/storage/manager/reform"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

func main() {
	// config.
	cfg := config.Parse()

	// log system.
	log, err := logger.CreateZapLogger(cfg.LogLevel)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to create logger: %v", err)
	}
	defer log.Sync()

	ctxWithCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	databaseConn, err := db.CreateConnection(ctxWithCancel, cfg, log)
	if err != nil {
		log.Info("failed to connect to users database: %v", err)
	}

	reformConn := reform.NewDB(databaseConn.DB, postgresql.Dialect, reform.NewPrintfLogger(log.Printf))

	// authentication.
	users := usersStorage.Create(reformConn, log)
	jwtGenerator := jwtgenerator.Create(cfg.JWTKey, cfg.JWTExpiration, log)
	authController := auth.CreateController(users, jwtGenerator, log)

	// data.
	storageManager := reformManager.CreateReform(reformConn, log)
	dataStorage := storage.Create(storageManager, log)
	mainController := controller.Create(dataStorage, log)

	// controllers mounting.
	server := fiber.New()
	server.Use(log.LogHandler)

	server.Mount("/api/user", authController.Route())

	serverData := server.Group("/api")
	serverData.Use(authController.AuthorizeUser(data.RoleUser))
	serverData.Mount("/", mainController.Route())

	// server launch.
	go func(log logger.BaseLogger) {
		log.Info("server is launching with host '%s'", cfg.HostAddr)
		if err = server.Listen(cfg.HostAddr); err != nil {
			log.Info("failed to launch server: %v", err)
		}

		log.Info("server has been stopped")
	}(log)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

}
