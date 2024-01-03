package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/dummy-image/internal/config"
	"github.com/hixraid/dummy-image/internal/middleware"
)

func InitHandler(cfg *config.ImageConfig) http.Handler {
	router := gin.New()

	router.LoadHTMLGlob("templates/*")

	router.Static("/static", "./static")
	router.StaticFile("favicon.ico", "./resources/favicon.ico")

	router.GET("/", index)

	parserMiddleware := middleware.NewImageURLParser(cfg)
	router.NoRoute(parserMiddleware.ParseURL, generateImage)

	return router
}
