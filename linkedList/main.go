package main

import (
	"fmt"
)

// StationListフィールドを定義
type StationList struct {
	name string // 駅名
	next int // ポインタ
}

func main() {
	s := []struct {
		name string
		next int
	}{
		{"new Osaka", -1},
		{"Nagoya", 3},
		{"Tokyo", 4},
		{"Kyoto", 0},
		{"new Yokohama", 1},
	}

	for i := 2; i != -1; {
		fmt.Print("[ ", s[i].name, " ]->")
		i = s[i].next
	}
}