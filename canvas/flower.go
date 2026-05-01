package canvas

import (
	"image"
	"math"
)

func DrawFlower(img *image.RGBA, steps []int, cx, cy int) {
	total := len(steps)
	petals := 16
	angleStep := 2 * math.Pi / float64(petals)

	for p := 0; p < petals; p++ {
		baseAngle := float64(p) * angleStep

		for i := 0; i < total; i++ {
			if steps[i] == 0 {
				break
			}
			t := float64(i) / float64(total)

			// Petalform: Rose curve — r = cos(k*theta)
			theta := t * math.Pi
			r := float64(steps[i]) * 2.0 * math.Abs(math.Cos(3*theta))

			angle := baseAngle + theta
			c := StepColor(p*total/petals+i, total)

			x1 := cx + int(r*0.8*math.Cos(angle))
			y1 := cy + int(r*0.8*math.Sin(angle))
			x2 := cx + int(r*math.Cos(angle+0.05))
			y2 := cy + int(r*math.Sin(angle+0.05))

			DrawThickLine(img, x1, y1, x2, y2, 2, c)
		}
	}
}
