package main

import (
	"fmt"
)

type BST struct {
	date  int
	left  int
	right int
}

var rootIdx, newIdx int
var tree [10]BST

func main() {
	addBST(4)
	addBST(6)
	addBST(5)
	addBST(2)
	addBST(3)
	addBST(7)
	addBST(1)
	printPhysicalBST()

	// 論理的な順序で出力
	fmt.Println("----------------------------")
	printLogicalBST(rootIdx)

	// 二分探索木を探索
	fmt.Println("'5'の探索結果 =", searchBST(5, rootIdx))
	fmt.Println("'8'の探索結果 =", searchBST(8, rootIdx))
}

func addBST(date int) {
	var addFlag bool

	// 物理的な位置に追加
	tree[newIdx] = BST{date, -1, -1}

	// 根のデータないなら、理論的な位置にポインタを設定
	if newIdx != rootIdx {
		currentIdx := rootIdx

		for addFlag == false {
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
				if tree[currentIdx].right == -1 {
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

func printPhysicalBST() {
	for i := 0; i < newIdx; i++ {
		fmt.Printf(
			"tree[%d]:data = %d, left = %d, right = %d\n",
			i,
			tree[i].date,
			tree[i].left,
			tree[i].right,
		)
	}
}

func printLogicalBST(currentIdx int) {
	if currentIdx != -1 {
		fmt.Printf(
			"tree[%d]:data = %d, left = %d, right = %d\n",
			currentIdx,
			tree[currentIdx].date,
			tree[currentIdx].left,
			tree[currentIdx].right,
		)

		// 再帰呼び出し
		printLogicalBST(tree[currentIdx].left)
		printLogicalBST(tree[currentIdx].right)
	}
}

func searchBST(x int, currentIdx int) int {
	if currentIdx == -1 {
		return -1
	}

	if tree[currentIdx].date == x {
		return currentIdx
	}

	if tree[currentIdx].date > x {
		return searchBST(x, tree[currentIdx].left)
	}

	return searchBST(x, tree[currentIdx].right)
}
