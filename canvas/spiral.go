package canvas

import (
	"image"
	"math"
)

// DrawSpiral draws a rectangular Euclidean spiral based on GCD steps
func DrawSpiral(img *image.RGBA, steps []int, cx, cy int) {
	x, y := cx, cy
	total := len(steps)

	for i := 0; i < total-1; i++ {
		length := steps[i]
		if length == 0 {
			break
		}
		c := StepColor(i, total)
		angle := float64(i) * math.Pi / 4

		x2 := x + int(float64(length)*math.Cos(angle))
		y2 := y + int(float64(length)*math.Sin(angle))

		DrawLine(img, x, y, x2, y2, c)
		x, y = x2, y2
	}
}
