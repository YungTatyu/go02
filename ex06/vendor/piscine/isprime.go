package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}

	limit := Sqrt(nb) + 1
	for i := 2; i < limit; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
