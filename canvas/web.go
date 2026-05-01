package canvas

import (
	"image"
	"math"
)

func DrawWeb(img *image.RGBA, steps []int, cx, cy int) {
	total := len(steps)
	bounds := img.Bounds()
	radius := float64(min(bounds.Max.X, bounds.Max.Y)/2) * 0.9

	spokes := total
	if spokes > 360 {
		spokes = 360
	}
	angleStep := 2 * math.Pi / float64(spokes)

	spokeX := make([]int, spokes)
	spokeY := make([]int, spokes)

	// Rita spekar med unik färg per spake
	for i := 0; i < spokes; i++ {
		angle := float64(i) * angleStep
		spokeX[i] = cx + int(radius*math.Cos(angle))
		spokeY[i] = cy + int(radius*math.Sin(angle))
		c := StepColor(i, spokes)
		DrawThickLine(img, cx, cy, spokeX[i], spokeY[i], 1, c)
	}

	// Koppla spekar med färg baserad på position i hela steglistan
	for i := 0; i < spokes-1; i++ {
		if steps[i] == 0 {
			break
		}
		c := StepColor(i, spokes)
		jump := steps[i] % spokes
		if jump == 0 {
			jump = 1
		}
		target := (i + jump) % spokes
		DrawThickLine(img, spokeX[i], spokeY[i], spokeX[target], spokeY[target], 2, c)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
