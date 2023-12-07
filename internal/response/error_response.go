package response

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(ctx *gin.Context, status int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(status, errorResponse{message})
}
