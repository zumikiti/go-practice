package main

import (
	"fmt"
	"math/rand"
)

var (
	knapMax int = 6 // ナップザックの耐久量
	itemNum int = 5 // 品物の種類
	indNum int = 8 // 個体数の数
	mutateRate float32 = 0.1 // 突然変異する確率
	itemName = []string{"A", "B", "C", "D", "E"}
	itemWeight = []int{1, 2, 3, 4, 5}
	itemValue = []int{100, 300, 350, 500, 650}
	indGeneration int
	indGene = make([][]int, indNum)
	indWeight = make([]int, indNum)
	indValue = make([]int, indNum)
	indFitness = make([]int, indNum)
)

func main() {
	createIndividual()
	calcIndvidual()
	showIndividual()
	sortIndividual()
	showIndividual()
	selectIndividual()
	crossoverIndividual()
	mutateIndividual()
	calcIndvidual()
	sortIndividual()
	showIndividual()
}

// 個体をランダムに生成する
func createIndividual() {
	for ind := 0; ind < indNum; ind++ {
		for item := 0; item < itemNum; item++ {
			r := rand.Float32()
			i := 0
			if r > 0.5 {
				i = 1
			}
			indGene[ind] = append(indGene[ind], i)
		}
	}
}

// 個体の重量、価格、適用度を計算する
func calcIndvidual() {
	for ind := 0; ind < indNum; ind++ {
		// 重量と価値を計算する
		indWeight[ind] = 0
		indValue[ind] = 0
		for item := 0; item < itemNum; item++ {
			if indGene[ind][item] == 1 {
				indWeight[ind] += itemWeight[item]
				indValue[ind] += itemValue[item]
			}
		}

		// 適用度を計算する
		if indWeight[ind] <= knapMax {
			indFitness[ind] = indValue[ind]
		} else {
			indFitness[ind] = 0
		}
	}
}

// 個体の情報を表示する
func showIndividual() {
	// 世代を表示する
	fmt.Printf("\n＜第%d世代＞\n", indGeneration)

	// 遺伝子、重量、適用度を表示する
	fmt.Printf("遺伝子\t\t重量\t価値\t適用度\n")

	for ind := 0; ind < indNum; ind++ {
		for item := 0; item < itemNum; item++ {
			fmt.Printf("[%v]", indGene[ind][item])
		}
		fmt.Printf("\t%2dkg\t%4d円\t%4d\n", indWeight[ind], indValue[ind], indFitness[ind])
	}
	fmt.Print("\n")
}

// 個体の適用度の大きい順にソートする
func sortIndividual() {
	// 挿入法でソートする
	for pos := 1; pos < indNum; pos++ {
		ins := pos
		for ins >= 1 && indFitness[ins - 1] < indFitness[ins] {
			for item := 0; item < itemNum; item++ {
				tmp := indGene[ins - 1][item]
				indGene[ins - 1][item] = indGene[ins][item]
				indGene[ins][item] = tmp
			}

			tmp := indWeight[ins - 1]
			indWeight[ins - 1] = indWeight[ins]
			indWeight[ins] = tmp

			tmp = indValue[ins - 1]
			indValue[ins - 1] = indValue[ins]
			indValue[ins] = tmp

			tmp = indFitness[ins - 1]
			indFitness[ins - 1] = indFitness[ins]
			indFitness[ins] = tmp

			ins--
		}
	}
}

// 淘汰するメソット
func selectIndividual() {
	// 適用度の上位50％を下位50％にコピーする
	for ind := 0; ind < indNum / 2; ind++ {
		for item := 0; item < itemNum; item++ {
			indGene[ind + indNum / 2][item] = indGene[ind][item]
		}
	}
	fmt.Println("下位50％を淘汰しました")
}

// 交配するメソッド
func crossoverIndividual() {
	var crossoverPoint int

	 // 下位50％にコピーした個体を対象とする
	 for ind := indNum / 2; ind < (indNum - 1); ind += 2 {
		// 交配する位置をランダムに決める
		r := rand.Float32()
		crossoverPoint = int(r * 10000) % (itemNum - 1) + 1
		for item := crossoverPoint; item < itemNum; item++ {
			// 隣の個体と交配する
			tmp := indGene[ind][item]
			indGene[ind][item] = indGene[ind + 1][item]
			indGene[ind + 1][item] = tmp
		}

		fmt.Printf(
			"個体%dと個体%dを%dの位置で交配しました。\n",
			ind,
			ind + 1,
			crossoverPoint,
		)
	}
}

// 突然変異するメソッド
func mutateIndividual() {
	// 下位50％を対象にする
	for ind := indNum / 2; ind < indNum; ind++ {
		for item := 0; item < itemNum ; item++ {
			// あらかじめ決めた確率で突然変異する
			if rand.Float32() <= mutateRate {
				 // 反転する
				 indGene[ind][item] ^= 1
				 fmt.Printf(
				 	"個体%dの%dの位置で突然変異しました。\n",
				 	ind,
				 	item,
				 )
			}
		}
	}
}