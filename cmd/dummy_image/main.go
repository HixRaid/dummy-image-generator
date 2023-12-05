package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hixraid/dummy-image/internal/config"
	"github.com/hixraid/dummy-image/internal/handler"
	"github.com/hixraid/dummy-image/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	cfgFile, err := config.LoadConfig("config")
	if err != nil {
		logrus.Fatalf("can't load file: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		logrus.Fatalf("can't parse config file: %v", err)
	}

	handler := handler.InitHandler()

	srv := server.New(cfg.Addr, handler)
	go func() {
		if err := srv.Run(); err != nil {
			logrus.Fatalf("error: '%v' while running http server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := srv.Shutdown(); err != nil {
		logrus.Errorf("error: '%v' while closing http server", err)
	}
}
