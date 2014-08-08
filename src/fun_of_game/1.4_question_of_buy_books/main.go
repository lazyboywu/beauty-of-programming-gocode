package main

import (
	"fmt"
	"fun_of_game/1.4_question_of_buy_books/algorithm_1"
)

func main() {
	var (
		X1 = 2
		X2 = 2
		X3 = 1
		X4 = 1
		X5 = 1
	)

	CalculateDiscount := algorithm_1.CalculateDiscount(X1, X2, X3, X4, X5)
	fmt.Printf("X1:%d,X2:%d,X3:%d,X4:%d,X5:%d CalculateDiscount=%f", X1, X2, X3, X4, X5, CalculateDiscount)
}
