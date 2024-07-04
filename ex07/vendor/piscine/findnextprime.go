package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	var primeNum int = 2
	for i := nb; ; i++ {

		if IsPrime(i) {
			primeNum = i
			break
		}
		if i == IntMax {
			break
		}
	}
	return primeNum
}
