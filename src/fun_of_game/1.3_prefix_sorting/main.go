package main

import (
	//"fmt"
	// 顺便学习下自定义包
	"fun_of_game/1.3_prefix_sorting/sorting"
)

func main() {

	s := sorting.NewPrefixSorting()

	s.Run([]int{3, 2, 1, 6, 5, 4, 9, 8, 7, 0}, 10)

	s.Output()

}
