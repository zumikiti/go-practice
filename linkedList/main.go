package main

import (
	"fmt"
)

// stationListフィールドを定義
type stationList []struct {
	name string // 駅名
	next int // ポインタ
}

var (
	s = stationList{
		{"Shin Osaka", -1},
		{"Nagoya", 3},
		{"Tokyo", 4},
		{"Kyoto", 0},
		{"Shin Yokohama", 1},
	}
)

func main() {
	// 初期値の連結リスト
	printstationList(s)

	// 連結リストに要素を追加
	insertSstationList(5, "Shinagawa", 2)
	printstationList(s)
}

func printstationList(l stationList) {
	for i := 2; i != -1; {
		fmt.Print("[ ", l[i].name, " ]->")
		i = l[i].next
	}
	fmt.Println()
}

func insertSstationList(insId int, insName string, prevId int) {
	insertObject := stationList{
		{insName, s[prevId].next},
	}
	s[prevId].next = insId
	s = append(s, insertObject...)
}