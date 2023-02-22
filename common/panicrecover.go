package main

import "fmt"

func divide(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered:", err)
		}
	}()

	if b == 0 {
		panic("Cannot divide by zero")
	}

	return a / b
}

func main() {
	result := divide(10, 2)
	fmt.Println(result)

	result = divide(10, 0)
	fmt.Println(result)
}
