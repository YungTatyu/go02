package piscine

const (
	UintMax = ^uint(0)
	IntMax  = int(UintMax >> 1)
)

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	re := RecursiveFactorial(nb - 1)
	if re > (IntMax / nb) {
		return 0
	}
	return nb * re
}
