package main

import (
	"fmt"
)

func main() {
	for n := 0; n <= 8; n++ {
		fmt.Print(fibonacci(n), ", ")
	}
	fmt.Print("\n")
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
