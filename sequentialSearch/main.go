package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := [10]int{86, 12, 9, 65, 87, 96, 36, 2, 8, 11}
	pos := -1
	x, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < len(a) && pos == -1; i++ {
		if a[i] == x {
			pos = i
		}
	}

	fmt.Printf("pos = %v", pos)
}
