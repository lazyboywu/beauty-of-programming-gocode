package algorithm_1

import (
	"math"
	"sort"
)

func CalculateDiscount(Y1 int, Y2 int, Y3 int, Y4 int, Y5 int) float64 {

	if Y1 == 0 && Y2 == 0 && Y3 == 0 && Y4 == 0 && Y5 == 0 {
		return 0
	}

	Y1, Y2, Y3, Y4, Y5 = TransformSort(Y1, Y2, Y3, Y4, Y5)
	var C1, C2, C3, C4, C5 float64
	var T1, T2, T3, T4, T5 int

	if Y5 >= 1 {
		T1, T2, T3, T4, T5 = TransformSort(Y1-1, Y2-1, Y3-1, Y4-1, Y5-1)
		C1 = 5*8*(1-0.25) + CalculateDiscount(T1, T2, T3, T4, T5)
	} else {
		C1 = 0
	}

	if Y4 >= 1 {
		T1, T2, T3, T4, T5 = TransformSort(Y1-1, Y2-1, Y3-1, Y4-1, Y5)
		C2 = 4*8*(1-0.20) + CalculateDiscount(T1, T2, T3, T4, T5)
	} else {
		C2 = 0
	}

	if Y3 >= 1 {
		T1, T2, T3, T4, T5 = TransformSort(Y1-1, Y2-1, Y3-1, Y4, Y5)
		C3 = 3*8*(1-0.10) + CalculateDiscount(T1, T2, T3, T4, T5)
	} else {
		C3 = 0
	}

	if Y2 >= 1 {
		T1, T2, T3, T4, T5 = TransformSort(Y1-1, Y2-1, Y3, Y4, Y5)
		C4 = 2*8*(1-0.05) + CalculateDiscount(T1, T2, T3, T4, T5)
	} else {
		C4 = 0
	}

	if Y1 >= 1 {
		T1, T2, T3, T4, T5 = TransformSort(Y1-1, Y2, Y3, Y4, Y5)
		C5 = 8 + CalculateDiscount(T1, T2, T3, T4, T5)
	} else {
		C5 = 0
	}

	return Min(C1, C2, C3, C4, C5)
}

func TransformSort(Y1 int, Y2 int, Y3 int, Y4 int, Y5 int) (int, int, int, int, int) {
	data := []int{Y1, Y2, Y3, Y4, Y5}
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	return data[0], data[1], data[2], data[3], data[4]
}

func Min(C1 float64, C2 float64, C3 float64, C4 float64, C5 float64) float64 {
	return min(C1,
		min(C2,
			min(C3,
				min(C4, C5),
			),
		),
	)
}

func min(a float64, b float64) float64 {
	if a == 0 && b == 0 {
		return 0
	}
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return math.Min(a, b)
}
