package main

import (
	"fmt"
)

var (
	// ナップザックの耐久量
	knapMax int = 6
	// 品物の種類
	itemNum int = 5
	// 品物の名称
	name = []string{"A", "B", "C", "D", "E"}
	// 品物の重量
	weight = []int{1, 2, 3, 4, 5}
	// 品物の価値
	value = []int{100, 300, 350, 500, 650}
	// 品物を吟味した直後の価値
	maxVlaue = make([][]int, itemNum, knapMax)
	// 最後に入れた品物
	lastItem = make([]int, knapMax+1)
)

func main() {
	tasteKanpZero()

	testeKanp()

	answer()
}

func showKnap(item int) {
	var knap int // 0-6kg knap

	// 吟味した結果の情報を表示する
	fmt.Printf("<%v, %dkg, %d円を吟味した結果>\n", name[item], weight[item], value[item])

	// ナップザックの耐久量を表示する
	for knap = 0; knap <= knapMax; knap++ {
		fmt.Printf("%dkg\t", knap)
	}
	fmt.Print("\n")

	// ナップザックに詰められた品物の価値の合計を表示する
	for knap = 0; knap <= knapMax; knap++ {
		fmt.Printf("%d円\t", maxVlaue[item][knap])
	}
	fmt.Print("\n")

	// ナップザックへ最後に入れたものを表示する
	for knap = 0; knap <= knapMax; knap++ {
		if lastItem[knap] != -1 {
			fmt.Printf("%v\t", name[lastItem[knap]])
		} else {
			fmt.Print("なし\t")
		}
	}

	fmt.Println("\n\n")
}

// 0番目の品物を吟味する
func tasteKanpZero() {
	item := 0

	for knap := 0; knap <= knapMax; knap++ {
		// 耐重量以下なら選ぶ
		if weight[item] <= knap {
			maxVlaue[item] = append(maxVlaue[item], value[item])
			lastItem[knap] = item
		} else {
			// 耐久量以下でないなら選ばない
			maxVlaue[0] = append(maxVlaue[0], 0)
			lastItem[knap] = -1
		}
	}
	showKnap(item)
}

// 1番目〜itemNum -1番目の品物まで吟味
func testeKanp() {
	for item := 1; item < itemNum; item++ {
		for knap := 0; knap <= knapMax; knap++ {
			// 耐重量以下なら選ぶ
			if weight[item] <= knap {
				// 選んだ場合の価格を求める
				selVal := maxVlaue[item-1][knap-weight[item]] + value[item]
				// 価格が大きくなるなら、選ぶ
				if selVal > maxVlaue[item-1][knap] {
					maxVlaue[item] = append(maxVlaue[item], selVal)
					lastItem[knap] = item
				} else {
					// 価格が大きくならないなら、選ばない
					maxVlaue[item] = append(maxVlaue[item], maxVlaue[item-1][knap])
				}
			} else {
				// 耐久量以下でないなら選ばない
				maxVlaue[item] = append(maxVlaue[item], maxVlaue[item-1][knap])
			}
		}
		showKnap(item)
	}
}

func answer() {
	fmt.Println("＜ナップザックに入っている商品を調べる＞")
	var totalWeight, item int
	for kanp := knapMax; kanp > 0; kanp -= weight[item] {
		item = lastItem[kanp]
		fmt.Printf(
			"%dkgのナップザックに最後に入れた品物は%vです。\n",
			kanp,
			name[item],
		)
		totalWeight += weight[item]
		fmt.Printf(
			" %v, %dkg, %d円\n",
			name[item],
			weight[item],
			value[item],
		)
		fmt.Printf(
			" %dkg - %dkg = %dkgです。\n",
			kanp,
			weight[item],
			kanp-weight[item],
		)
	}
	fmt.Println("\n<解を表示する>")
	fmt.Printf("重量の合計値　＝　%dkg\n", totalWeight)
	fmt.Printf("価値の最大値　＝　%d円\n", maxVlaue[itemNum-1][knapMax])
}
