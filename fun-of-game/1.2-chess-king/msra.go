package main

import (
	"fmt"
)

func main() {
	i := 81
	// for i-- { 不可以
	for i > 0 {
		i--
		if i/9%3 == i%9%3 {
			continue
		}
		fmt.Printf("A = %d, B = %d\n", i/9+1, i%9+1)
	}
}
