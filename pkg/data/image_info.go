package data

import "image/color"

type ImageFormat uint8

const (
	PNGFormat  ImageFormat = 1
	JPEGFormat ImageFormat = 2
	SVGFormat  ImageFormat = 3
)

type ImageInfo struct {
	Size            [2]int
	BackgroundColor color.Color
	TextColor       color.Color
	Format          ImageFormat
	Text            string
}
