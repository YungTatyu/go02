package piscine

const (
	UintMax = ^uint(0)
	IntMax  = int(UintMax >> 1)
)

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	re := RecursivePower(nb, power-1)
	return nb * re
}
