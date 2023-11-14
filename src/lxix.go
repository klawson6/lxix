package lxix

func RomanNumerals(num int) (rn string) {
	if num < 0 || num > 3999 {
		panic("Input out of range")
	}
	for i := 0; i < num/1000; i++ {
		rn += "M"
	}
	// oh dear

	return
}

func convertDigit(digit int, uc string, lc string) (s string) {
	// wtf do
	return
}
