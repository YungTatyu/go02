package piscine

const (
	UintMax = ^uint(0)
	IntMax  = int(UintMax >> 1)
)

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	for i := 1; i <= (nb/2 + 1); i++ {
		if i > (IntMax / i) {
			return 0
		}
		square := i * i
		if square == nb {
			return i
		}
		if square > nb {
			break
		}
	}
	return 0
}
