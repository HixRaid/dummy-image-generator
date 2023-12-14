package data

import "image/color"

type ImageFormat uint8

const (
	PNG  ImageFormat = 1
	JPEG ImageFormat = 2
	SVG  ImageFormat = 3
)

type ImageResolution [2]int

var (
	QVGA  ImageResolution = [2]int{320, 240}
	SIF   ImageResolution = [2]int{352, 240}
	CIF   ImageResolution = [2]int{352, 288}
	WQVGA ImageResolution = [2]int{400, 240}
	HVGA  ImageResolution = [2]int{640, 240}
	NHD   ImageResolution = [2]int{640, 360}
	VGA   ImageResolution = [2]int{640, 480}
	WVGA  ImageResolution = [2]int{800, 480}
	SVGA  ImageResolution = [2]int{800, 600}
	FWVGA ImageResolution = [2]int{848, 480}
	QHD   ImageResolution = [2]int{960, 540}
	WSVGA ImageResolution = [2]int{1024, 600}
	XGA   ImageResolution = [2]int{1024, 768}
	WXVGA ImageResolution = [2]int{1200, 600}
	HD    ImageResolution = [2]int{1280, 720}
	WXGA  ImageResolution = [2]int{1280, 768}
	SXGA  ImageResolution = [2]int{1280, 1024}
	XJXGA ImageResolution = [2]int{1536, 960}
	WSXGA ImageResolution = [2]int{1600, 1024}
	UXGA  ImageResolution = [2]int{1600, 1200}
	FHD   ImageResolution = [2]int{1920, 1080}
	WUXGA ImageResolution = [2]int{1920, 1200}
)

type ImageInfo struct {
	Size            ImageResolution
	BackgroundColor color.Color
	TextColor       color.Color
	Text            string
}
