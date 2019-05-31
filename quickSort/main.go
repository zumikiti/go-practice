package main

import (
	"fmt"
)

func main() {
	a := []int{4, 7, 1, 6, 2, 5, 3}

	printArray(a)

	sortArray(a, 0, len(a) - 1)

	printArray(a)
}

func printArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("[%v]", a[i])
	}
	fmt.Println("\n")
}

 // 配列a[start] - a[end]を昇順にソートする
func sortArray(a []int, start int, end int) {
	var pivot int

	// 配列が二つ以上ある場合
	if start < end {
		pivot = divideArray(a, start, end)
		fmt.Println("pivot", pivot)
		sortArray(a, start, pivot - 1)

		sortArray(a, pivot + 1, end)
	}
}

func divideArray(a []int, head int, tail int) int {
	var left, right, temp int
	left = head + 1 // 先頭+1からたどる位置
	right = tail   // 末端からたどる位置

	// 基準値a[head]より小さい要素を前側に、大きい要素を後ろ側に移動する
	for {
		fmt.Println(a)
		// 配列を先頭+1から後ろに向かってたどり、基準値より大きい要素を見つける
		for left < tail && a[head] > a[left] {
			left++
		}
		// 配列の末尾から前に向かってたどり、基準値より小さい要素を見つける
		for a[head] < a[right] {
			right--
		}

		// チェックする要素がなくなったら終了する
		if left >= right {
			break
		}

		// 基準値より大きいa[left]と、より小さいa[right]を交換する
		temp = a[left]
		a[left] = a[right]
		a[right] = temp

		fmt.Printf("temp=%v, a[left]=%v, a[right]=%v\n", temp, a[left], a[right])
		// 次の要素のチェックに進む
		left++
		right--
	}

	// 基準値a[head]とa[right]を交換する
	temp = a[head]
	a[head] = a[right]
	a[right] = temp

	// 基準値a[right]の位置を返す
	return a[right]
}