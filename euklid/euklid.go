package euklid

import (
	"math"
	"math/big"
)

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// TruncateBig returnerar de första n siffrorna av x som ett nytt big.Int.
func TruncateBig(x *big.Int, digits int) *big.Int {
	s := x.String()
	if len(s) > digits {
		s = s[:digits]
	}
	n := new(big.Int)
	n.SetString(s, 10)
	return n
}

// StepsBig returnerar normaliserade steg (1–maxStep) med logaritmisk skala.
// Logaritmisk skala behövs eftersom primtal ger extremt ojämn fördelning —
// ett par enorma steg och tusentals winziga — linjär skala gör alla små steg = 1.
func StepsBig(a, b *big.Int, maxStep int) []int {
	zero := big.NewInt(0)
	var rawSteps []*big.Int

	a = new(big.Int).Set(a)
	b = new(big.Int).Set(b)

	for b.Cmp(zero) != 0 {
		q := new(big.Int)
		q.Div(a, b)
		rawSteps = append(rawSteps, new(big.Int).Set(q))
		a, b = b, new(big.Int).Mod(a, b)
	}

	// Konvertera till float64 via log för att hantera extrema värden
	logVals := make([]float64, len(rawSteps))
	maxLog := 0.0
	for i, q := range rawSteps {
		f, _ := new(big.Float).SetInt(q).Float64()
		if f < 1 {
			f = 1
		}
		l := math.Log(f)
		logVals[i] = l
		if l > maxLog {
			maxLog = l
		}
	}

	steps := make([]int, len(rawSteps))
	for i, l := range logVals {
		if maxLog == 0 {
			steps[i] = 1
			continue
		}
		v := int(l / maxLog * float64(maxStep))
		if v < 1 {
			v = 1
		}
		steps[i] = v
	}
	return steps
}
