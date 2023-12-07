package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/dummy-image/internal/middleware"
)

func InitHandler() http.Handler {
	router := gin.New()

	router.NoRoute(middleware.SplitPath, generateImage)

	return router
}
