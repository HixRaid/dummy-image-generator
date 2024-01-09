package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
