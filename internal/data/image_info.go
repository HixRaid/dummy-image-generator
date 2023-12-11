package data

import "image/color"

type ImageInfo struct {
	Size            [2]int
	BackgroundColor color.Color
	TextColor       color.Color
	Format          string
	Text            string
}
