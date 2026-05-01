package canvas

import (
	"image"
	"math"
)

func DrawWave(img *image.RGBA, steps []int, cx, cy int) {
	total := len(steps)
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	layers := 20
	if layers > total {
		layers = total
	}

	for layer := 0; layer < layers; layer++ {
		amplitude := float64(steps[layer%total]) * 0.8
		if amplitude > float64(height)/4 {
			amplitude = float64(height) / 4
		}
		frequency := float64(layer+1) * 0.5
		yBase := float64(cy) + float64(layer-layers/2)*float64(height/layers)
		c := StepColor(layer, layers)

		prevX := 0
		prevY := int(yBase + amplitude*math.Sin(0))

		for x := 1; x < width; x++ {
			t := float64(x) / float64(width) * 2 * math.Pi * frequency
			y := int(yBase + amplitude*math.Sin(t))
			DrawThickLine(img, prevX, prevY, x, y, 2, c)
			prevX = x
			prevY = y
		}
	}
}
