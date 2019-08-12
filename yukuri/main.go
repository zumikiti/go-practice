package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 引数がなければエラー
	c := len(os.Args) - 1
	if c < 2 {
		fmt.Fprintln(os.Stderr, "引数が整数で二つ以上必要です。example: ./yukuri foo bar")
		os.Exit(1)
	}

	a, err := strconv.Atoi(os.Args[1])
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}

	result := a
	fmt.Println(result)
}
