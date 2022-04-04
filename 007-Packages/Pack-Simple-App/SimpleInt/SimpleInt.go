package SimpleInt

func SimpleInterest(p, r, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
