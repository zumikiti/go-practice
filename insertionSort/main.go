package main

import (
	"fmt"
)

func main() {
	arr := [6]int{78, 23, 123, 55, 33, 99}
	var i, c, temp int

	// ソート前の配列表示
	fmt.Println(arr)

	// 昇順でソート
	for i = 1; i < len(arr); i++ {
		temp = arr[i]
		for c = i - 1; c >= 0; c-- {
			if arr[c] > temp {
				arr[c+1] = arr[c]
			} else {
				break
			}
		}
		arr[c+1] = temp
	}

	// ソート結果を表示
	fmt.Println(arr)
}
