package main

import (
	"fmt"
)

type BST struct {
	date int
	left int
	right int
}

var rootIdx, newIdx int
var tree [10]BST

func main() {
	addBST(4)
	fmt.Println(tree)
}

func addBST(date int) {
	var addFlag bool

	// 物理的な位置に追加
	tree[newIdx] = BST{date, -1, -1}
fmt.Println(tree[newIdx])

	// 根のデータないなら、理論的な位置にポインタを設定
	if newIdx != rootIdx {
		currentIdx := rootIdx

		for ok := false; ok; ok = addFlag == true {
			// より小さい場合は、左を探る
			if date < tree[currentIdx].date {
				// 左が末端なら、そこに追加する
				if tree[currentIdx].left == -1 {
					tree[currentIdx].left = newIdx
					addFlag = true
				} else {
					currentIdx = tree[currentIdx].left
				}
			} else {
				// より大きい場合は、右に探る
				// 右が末端なら、そこに追加する
				if tree[currentIdx].left == -1 {
					tree[currentIdx].right = newIdx
					addFlag = true
				} else {
					currentIdx = tree[currentIdx].right
				}
			}
		}
	}

	newIdx++
}