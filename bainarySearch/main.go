package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := makeArr()
	rand.Seed(time.Now().UnixNano())
	var x uint64 = arr[rand.Intn(len(arr))]
	var pos, left, right int = -1, 0, len(arr) - 1

	for pos == -1 && left <= right {
		middle := ( left + right ) /2

		if arr[middle] == x {
			pos = middle
		}
		if arr[middle] > x {
			right = middle - 1
		}
		if arr[middle] < x {
			left = middle + 1
		}
	}

	fmt.Println(rand.Intn(len(arr)))
	fmt.Println(x)
	fmt.Print(pos)
}

func makeArr() [1000]uint64 {
	var arr [1000]uint64
	var a, b, sum uint64 = 1, 1, 0

	for i := 0; i < 1000; i++ {
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