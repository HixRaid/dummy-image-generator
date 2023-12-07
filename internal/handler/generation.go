package handler

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hixraid/dummy-image/internal/middleware"
	"github.com/hixraid/dummy-image/internal/response"
)

func generateImage(ctx *gin.Context) {
	path, err := middleware.GetPath(ctx)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	size, err := strconv.ParseInt(path[0], 10, 64)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, "invalid image size")
		return
	}

	var sizeInt = int(size)

	img := image.NewRGBA(image.Rect(0, 0, sizeInt, sizeInt))

	for x := 0; x < sizeInt; x++ {
		for y := 0; y < sizeInt; y++ {
			img.Set(x, y, color.Black)
		}
	}

	ctx.Stream(func(w io.Writer) bool {
		png.Encode(w, img)

		return false
	})
}
