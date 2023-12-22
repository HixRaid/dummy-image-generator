package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
