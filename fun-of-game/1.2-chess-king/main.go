package main

import (
	"fmt"
)

// 貌似当前版本不支持宏，那么就以常量和函数代替吧

const (
	HALF_BITS_LENGTH uint8 = 4
	FULLMASK         uint8 = 255
	GRIDW            uint8 = 3
)

const (
	// golang 常量赋值时，不能溢出，不能丢值
	//LMASK uint8 = (FULLMASK << HALF_BITS_LENGTH)
	LMASK uint8 = 0xf0
	//RMASK uint8 = (FULLMASK >> HALF_BITS_LENGTH)
	RMASK uint8 = 0x0f
)

func RSET(b *uint8, n uint8) {
	*b = (LMASK & *b) ^ n
}

func LSET(b *uint8, n uint8) {
	*b = (RMASK & *b) ^ (n << HALF_BITS_LENGTH)
}

func RGET(b uint8) uint8 {
	return (RMASK & b)
}

func LGET(b uint8) uint8 {
	return ((LMASK & b) >> HALF_BITS_LENGTH)
}

func main() {
	var b uint8
	LSET(&b, 1)
	LSET(&b, 2)

	for LSET(&b, 1); LGET(b) <= GRIDW*GRIDW; LSET(&b, (LGET(b) + 1)) {
		for RSET(&b, 1); RGET(b) <= GRIDW*GRIDW; RSET(&b, (RGET(b) + 1)) {
			if LGET(b)%GRIDW != RGET(b)%GRIDW {
				fmt.Printf("A = %d, B = %d\n", LGET(b), RGET(b))
			}
		}
	}
}
