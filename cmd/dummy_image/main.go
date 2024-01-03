package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/dummy-image/internal/config"
	"github.com/hixraid/dummy-image/internal/handler"
	"github.com/hixraid/dummy-image/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	cfgFile, err := config.LoadConfig("config")
	if err != nil {
		logrus.Fatalf("can't load file: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		logrus.Fatalf("can't parse config file: %v", err)
	}

	fmt.Println(cfg)
	fmt.Println(cfg.Server)
	fmt.Println(cfg.Image)
	fmt.Println(cfg.Image.Size)
	fmt.Println(cfg.Image.Color)

	handler := handler.InitHandler(cfg.Image)

	srv := server.New(cfg.Server.Addr, handler)
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
