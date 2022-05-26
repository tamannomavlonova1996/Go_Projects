package transfer

func Total(amount float64) (sum float64) {
	persent := 0.5
	sum = (amount * persent / 100) + amount
	return sum
}
