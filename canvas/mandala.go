package canvas

import (
	"image"
	"math"
)

func DrawMandala(img *image.RGBA, steps []int, cx, cy, symmetry int) {
	total := len(steps)
	angleStep := 2 * math.Pi / float64(symmetry)

	for s := 0; s < symmetry; s++ {
		baseAngle := float64(s) * angleStep
		r := 0.0

		for i := 0; i < total-1; i++ {
			length := float64(steps[i])
			if length == 0 {
				break
			}
			c := StepColor(i, total)
			angle := baseAngle + float64(i)*math.Pi/float64(total)

			x1 := cx + int(r*math.Cos(angle))
			y1 := cy + int(r*math.Sin(angle))
			r += length * 0.5
			x2 := cx + int(r*math.Cos(angle))
			y2 := cy + int(r*math.Sin(angle))

			DrawThickLine(img, x1, y1, x2, y2, 3, c)
		}
	}
}
