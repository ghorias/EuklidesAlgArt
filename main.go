package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"EuklidesAlgArt/canvas"
	"EuklidesAlgArt/euklid"
)

func loadPrime(filename string, maxDigits int) *big.Int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s := strings.TrimSpace(string(data))
	if len(s) > maxDigits {
		s = s[:maxDigits]
	}
	n := new(big.Int)
	n.SetString(s, 10)
	return n
}

func main() {
	fmt.Println("Läser primtal...")
	p1 := loadPrime("p2.txt", 10000)
	p2 := loadPrime("p1.txt", 10000)

	fmt.Println("Kör Euklides algoritm...")
	steps := euklid.StepsBig(p1, p2, 400)
	fmt.Printf("Antal steg: %d\n", len(steps))

	width, height := 2000, 2000
	cx, cy := width/2, height/2

	fmt.Println("Ritar mandala...")
	img1 := canvas.NewCanvas(width, height)
	canvas.DrawMandala(img1, steps, cx, cy, 12)
	if err := canvas.SaveCanvas(img1, "mandala.png"); err != nil {
		panic(err)
	}
	fmt.Println("Sparad: mandala.png")

	fmt.Println("Ritar spiral...")
	img2 := canvas.NewCanvas(width, height)
	canvas.DrawSpiral(img2, steps, cx, cy)
	if err := canvas.SaveCanvas(img2, "spiral.png"); err != nil {
		panic(err)
	}
	fmt.Println("Sparad: spiral.png")

	fmt.Println("Ritar blomma...")
	img3 := canvas.NewCanvas(width, height)
	canvas.DrawFlower(img3, steps, cx, cy)
	if err := canvas.SaveCanvas(img3, "flower.png"); err != nil {
		panic(err)
	}
	fmt.Println("Sparad: flower.png")

	fmt.Println("Ritar webb...")
	img4 := canvas.NewCanvas(width, height)
	canvas.DrawWeb(img4, steps, cx, cy)
	if err := canvas.SaveCanvas(img4, "web.png"); err != nil {
		panic(err)
	}
	fmt.Println("Sparad: web.png")

	fmt.Println("Ritar vågor...")
	img5 := canvas.NewCanvas(width, height)
	canvas.DrawWave(img5, steps, cx, cy)
	if err := canvas.SaveCanvas(img5, "wave.png"); err != nil {
		panic(err)
	}
	fmt.Println("Sparad: wave.png")
}
