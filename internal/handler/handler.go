package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/dummy-image/internal/middleware"
)

func InitHandler() http.Handler {
	router := gin.New()

	router.NoRoute(middleware.ParseURL, generateImage)

	return router
}
