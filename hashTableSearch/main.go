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
	// 格納データを入力する　
	inputHashTable()

	// ハッシュ表からサーチ
	searchHashTable()
}

func inputHashTable() {
	fmt.Println(hashTable)

	fmt.Printf("\n格納する値＝")
	data := scnInt()

	// マイナスの値だった場合、格納をやめる
	if data < 0 {
		return
	}

	// ハッシュ値を求める
	hashV := hashFunc(data)

	// データの格納位置を決める
	pos := hashV
	for hashTable[pos] != -1 {
		// 格納位置を一つ進める
		pos++

		// 末尾を超えたら先頭に戻す
		if pos >= len(hashTable) {
			pos = 0
		}

		// ハッシュ値の位置まで戻ったら、ハッシュ表が一杯なので、繰り返しを終了する
		if pos == hashV {
			return
		}
	}

	// ハッシ表が一杯の場合、MSGを表示して抜ける
	if hashTable[pos] != -1 {
		fmt.Println("ハッシュ表が一杯です")
	}

	// ハッシュ表が一杯でなければ、データを格納する
	hashTable[pos] = data
	inputHashTable()
}

func searchHashTable() {
	// サーチする値を入力
	fmt.Printf("\nサーチする値＝")
	data := scnInt()

	// マイナス値だった場合、サーチをやめる
	if data < 0 {
		return
	}

	// ハッシュ値を求める
	hashV := hashFunc(data)

	// データをサーチする。
	pos := hashV
	for hashTable[pos] != -1 && hashTable[pos] != data {
		// 探索位置を加ええる
		pos++

		// 末端を超えたら、先頭に戻す
		if pos >= len(hashTable) {
			pos = 0
		}

		// -1を見つけるか、ハッシュ値の位置まで戻ったらデータが見つからないことが確定なので、繰り返しを終了
		if hashTable[pos] == -1 || pos == hashV {
			break
		}
	}

	// サーチした結果を表示する
	if hashTable[pos] == data {
		fmt.Printf("%d番目に見つかりました\n", pos)
	} else {
		fmt.Print("見つかりませんでした")
	}

	searchHashTable()
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