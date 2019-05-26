package main

import (
	"fmt"
	"strconv"
)

var hashTable = [10]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

func hashFunc(i int) int {
	return i % 10
}

func main() {
	for {
		// 格納データを入力する　
		fmt.Printf("\n格納する値＝")
		data := scnInt()

		// マイナスの値だった場合、格納をやめる
		if data < 0 {
			break
		}

		// ハッシュ値を求める
		hashV := hashFunc(data)

		// ハッシュ表に格納する
		hashTable[hashV] = data
	}

	// ハッシュ表からサーチ
	for {
		// サーチする値を入力
		fmt.Printf("\nサーチする値＝")
		data := scnInt()

		// マイナス値だった場合、サーチをやめる
		if data < 0 {
			break
		}

		// ハッシュ値を求める
		hashV := hashFunc(data)

		// サーチした結果を表示する
		if hashTable[hashV] == data {
			fmt.Printf("%d番目に見つかりました\n", hashV)
		} else {
			fmt.Print("見つかりませんでした")
		}
	}
}

func scnInt() int {
	var s string
	fmt.Scanf("%s", &s)

	data, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}

	return data
}