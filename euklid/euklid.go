package euklid

import "math/big"

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

// StepsBig returnerar normaliserade steg (1–maxStep) för att kunna ritas.
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

	maxQ := big.NewInt(0)
	for _, q := range rawSteps {
		if q.Cmp(maxQ) > 0 {
			maxQ.Set(q)
		}
	}

	steps := make([]int, len(rawSteps))
	maxBig := big.NewInt(int64(maxStep))
	for i, q := range rawSteps {
		if maxQ.Sign() == 0 {
			steps[i] = 1
			continue
		}
		scaled := new(big.Int).Mul(q, maxBig)
		scaled.Div(scaled, maxQ)
		v := int(scaled.Int64())
		if v < 1 {
			v = 1
		}
		steps[i] = v
	}
	return steps
}
