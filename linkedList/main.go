package main

import (
	"fmt"
)

// stationListフィールドを定義
type stationList struct {
	name string // 駅名
	next int    // ポインタ
}

var s = []stationList{
	{"Shin Osaka", -1},
	{"Nagoya", 3},
	{"Tokyo", 4},
	{"Kyoto", 0},
	{"Shin Yokohama", 1},
}

func main() {
	// 初期値の連結リスト
	printStationList(s)

	// 連結リストに要素を追加
	insertSstationList(5, "Shinagawa", 2)
	printStationList(s)

	// 連結リストの要素を削除
	deleteStationList(5, 2)
	printStationList(s)
}

func printStationList(l []stationList) {
	for i := 2; i != -1; {
		fmt.Print("[ ", l[i].name, " ]->")
		i = l[i].next
	}
	fmt.Println()
}

func insertSstationList(insId int, insName string, prevId int) {
	insertObject := stationList{insName, s[prevId].next}
	s[prevId].next = insId
	s = append(s, insertObject)
}

func deleteStationList(deleId int, prevId int) {
	s[prevId].next = s[deleId].next
	s = append(s[:deleId], s[deleId+1:]...)
}
