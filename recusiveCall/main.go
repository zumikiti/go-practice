package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)
	i, _ := strconv.Atoi(s)
	ans := factornal(i)
	fmt.Println(ans)
}

func factornal(n int) int {
	if n == 0 {
		return 1
	}

	return n * factornal(n - 1) 
}