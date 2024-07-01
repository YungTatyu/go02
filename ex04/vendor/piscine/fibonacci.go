package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index < 2 {
		return index
	}
	return recurFib(index, 1, 0)
}

func recurFib(index, a, b int) int {
	if index == 0 {
		return b
	}
	return recurFib(index-1, b, a+b)
}
