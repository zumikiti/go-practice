package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := makeArr()
	rand.Seed(time.Now().Unix())
	var x int = arr[rand.Intn(len(arr))]
	var pos int = -1

	for i := 0; i < len(arr) && pos == -1; i++ {
		if x == arr[i] {
			pos = i
		}
	}

	fmt.Println("探す値は、", x)
	fmt.Print("配列の", pos, "番目にあります。")
}

func makeArr() [1000]int {
	var arr [1000]int
	var a, b, sum int = 1, 1, 0

	for i := 0; i < len(arr); i++ {
		if i < 2 {
			arr[i] = a
		} else {
			sum = a + b
			arr[i] = sum
			a = b
			b = sum
		}
	}

	return arr
}
