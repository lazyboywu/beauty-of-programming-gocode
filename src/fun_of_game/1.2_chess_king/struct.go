package main

import (
	"fmt"
)

func main() {
	var i struct {
		a uint8
		b uint8
	}

	for i.a = 1; i.a <= 9; i.a++ {
		for i.b = 1; i.b <= 9; i.b++ {
			if i.a%3 != i.b%3 {
				fmt.Printf("A = %d, B = %d\n", i.a, i.b)
			}
		}
	}
}
