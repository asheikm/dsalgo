package main

func main() {
	n := fibonacci(10)
}

// o(2^n)
func fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	return fibonacci(num-2) + fibonacci(num-1)
}
