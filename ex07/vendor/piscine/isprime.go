package piscine

const (
	UintMax = ^uint(0)
	IntMax  = int(UintMax >> 1)
)

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i <= (nb / i); i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
