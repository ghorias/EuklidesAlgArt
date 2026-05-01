package canvas

import (
	"image"
	"math"
)

func DrawFlower(img *image.RGBA, steps []int, cx, cy int) {
	total := len(steps)
	petals := 12
	angleStep := 2 * math.Pi / float64(petals)

	for p := 0; p < petals; p++ {
		baseAngle := float64(p) * angleStep
		r := 0.0

		for i := 0; i < total; i++ {
			if steps[i] == 0 {
				break
			}
			c := StepColor(i, total)
			t := float64(i) / float64(total)

			// Petalform: r växer och krymper per kronblad
			petalR := r * math.Abs(math.Sin(t*math.Pi))
			angle := baseAngle + t*math.Pi*0.5

			x1 := cx + int(petalR*math.Cos(angle))
			y1 := cy + int(petalR*math.Sin(angle))

			r += float64(steps[i]) * 0.3
			petalR2 := r * math.Abs(math.Sin(t*math.Pi))

			x2 := cx + int(petalR2*math.Cos(angle))
			y2 := cy + int(petalR2*math.Sin(angle))

			DrawThickLine(img, x1, y1, x2, y2, 3, c)
		}
		r = 0.0
	}
}
