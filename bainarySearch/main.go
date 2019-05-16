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
	var left, middle int
	var right int = len(arr) - 1

	for pos == -1 && left <= right {
		middle = ( left + right ) /2
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
	fmt.Println(arr)
	fmt.Println(x)
	fmt.Print(pos)
}

func makeArr() [1000]int {
	var arr [1000]int
	var a, b, sum int = 1, 1, 0

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

	shuffle(arr)

	return arr
}

func shuffle(data [1000]int) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}