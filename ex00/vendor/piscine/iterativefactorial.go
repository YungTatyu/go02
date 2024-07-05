package piscine

const (
	UintMax = ^uint(0)
	IntMax  = int(UintMax >> 1)
)

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	var re int = 1
	for i := 1; i <= nb; i++ {
		if re > (IntMax / i) {
			return 0
		}
		re *= i
	}
	return re
}
