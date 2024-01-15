package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/dummy-image-generator/internal/config"
	"github.com/hixraid/dummy-image-generator/internal/middleware"
)

func InitHandler(cfg *config.ImageConfig) http.Handler {
	router := gin.New()

	router.LoadHTMLGlob("views/*")

	router.Static("/static", "./static")
	router.StaticFile("favicon.ico", "./resources/favicon.ico")

	router.GET("/", indexPage)

	parserMiddleware := middleware.NewImageURLParser(cfg)
	router.NoRoute(parserMiddleware.ParseURL, generateImage)

	return router
}
