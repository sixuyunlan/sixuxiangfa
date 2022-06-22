package main

import "fmt"

type ValNode struct {
	row   int
	col   int
	value int
}

func main() {
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 3
	//
	//for _, v := range chessMap {
	//	for _, v2 := range v {
	//		fmt.Printf("%d\t", v2)
	//	}
	//	fmt.Println()
	//}
	var sparseArr []ValNode
	valNode := ValNode{
		11,
		11,
		0,
	}

	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode := ValNode{
					i,
					j,
					v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	//原始转稀疏

	for i, valNode := range sparseArr {
		fmt.Printf("%d:%d %d  %d\n", i, valNode.row, valNode.col, valNode.value)
	}
	//稀疏转原始
	var chessMap2 [11][11]int
	for i, valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.row][valNode.col] = valNode.value
		}
	}
	//
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
