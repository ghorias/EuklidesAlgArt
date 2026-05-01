package canvas

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func NewCanvas(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.SetRGBA(x, y, color.RGBA{0, 0, 0, 255})
		}
	}
	return img
}

func SaveCanvas(img *image.RGBA, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}

func DrawPixel(img *image.RGBA, x, y int, c color.RGBA) {
	img.SetRGBA(x, y, c)
}

func DrawLine(img *image.RGBA, x0, y0, x1, y1 int, c color.RGBA) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 > x1 {
		sx = -1
	}
	if y0 > y1 {
		sy = -1
	}
	err := dx - dy
	for {
		img.SetRGBA(x0, y0, c)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func DrawThickLine(img *image.RGBA, x0, y0, x1, y1, thickness int, c color.RGBA) {
	for t := -thickness / 2; t <= thickness/2; t++ {
		dx := abs(x1 - x0)
		dy := abs(y1 - y0)
		if dx > dy {
			DrawLine(img, x0, y0+t, x1, y1+t, c)
		} else {
			DrawLine(img, x0+t, y0, x1+t, y1, c)
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
