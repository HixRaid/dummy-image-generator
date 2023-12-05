package handler

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"strconv"
)

func InitHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var size int = 100

		if ok := r.URL.Query().Has("size"); ok {
			i, err := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "bad request")
				return
			}

			if i >= 1 && i <= 5000 {
				size = int(i)
			}
		}

		img := image.NewRGBA(image.Rect(0, 0, size, size))

		var color = color.RGBA{0, 0, 0, 255}

		for x := 0; x < size; x++ {
			for y := 0; y < size; y++ {
				img.Set(x, y, color)
			}
		}

		png.Encode(w, img)
	})

	return mux
}
