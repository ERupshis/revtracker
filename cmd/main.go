package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/erupshis/revtracker/internal/config"
	"github.com/erupshis/revtracker/internal/controller"
	"github.com/erupshis/revtracker/internal/db"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/erupshis/revtracker/internal/storage/manager"
	"github.com/go-chi/chi/v5"
)

func main() {
	//config.
	cfg := config.Parse()

	//log system.
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

	storageManager := manager.CreateReform(databaseConn, log)
	dataStorage := storage.Create(storageManager, log)
	mainController := controller.Create(dataStorage, log)

	//controllers mounting.
	router := chi.NewRouter()
	router.Use(log.LogHandler)

	router.Mount("/", mainController.Route())

	//server launch.
	go func() {
		log.Info("server is launching with Host setting: %s", cfg.HostAddr)
		if err := http.ListenAndServe(cfg.HostAddr, router); err != nil {
			log.Info("server refused to start with error: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

}
