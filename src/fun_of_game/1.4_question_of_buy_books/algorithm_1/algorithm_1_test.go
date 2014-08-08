package algorithm_1_test

import (
	"fun_of_game/1.4_question_of_buy_books/algorithm_1"
	"fun_of_game/common/test_helper"
	"testing"
)

var sortTestCases = []struct {
	params   []int
	expected []int
}{
	{[]int{1, 1, 1, 1, 0}, []int{1, 1, 1, 1, 0}},
	{[]int{1, 1, 1, 2, 2}, []int{2, 2, 1, 1, 1}},
	{[]int{1, 1, 2, 2, 2}, []int{2, 2, 2, 1, 1}},
}

func TestSort(t *testing.T) {
	for i, testCase := range sortTestCases {
		var result []int
		t1, t2, t3, t4, t5 := algorithm_1.TransformSort(
			testCase.params[0], testCase.params[1],
			testCase.params[2], testCase.params[3],
			testCase.params[4],
		)
		result = []int{t1, t2, t3, t4, t5}

		if !test_helper.IntArrayEquals(result, testCase.expected) {
			t.Errorf("#%d:\n got: %q\nwant: %q", i, result, testCase.expected)
		}
	}

}

var minTestCases = []struct {
	params   []float64
	expected float64
}{
	{[]float64{1, 2, 3, 4, 5}, 1},
	{[]float64{5, 4, 3, 2, 1}, 1},
	{[]float64{4, 2, 1, 5, 3}, 1},
}

func TestMin(t *testing.T) {
	for i, testCase := range minTestCases {
		result := algorithm_1.Min(
			testCase.params[0], testCase.params[1],
			testCase.params[2], testCase.params[3],
			testCase.params[4],
		)
		if result != testCase.expected {
			t.Errorf("#%d:\n got: %q\nwant: %q", i, result, testCase.expected)
		}
	}
}

var calculateTestCases = []struct {
	params   []int
	expected float64
}{
	{[]int{1, 0, 0, 0, 0}, 8},
	{[]int{1, 1, 0, 0, 0}, float64(2 * 8 * (1 - 0.05))},
}

func TestCalculate(t *testing.T) {
	for i, testCase := range calculateTestCases {
		result := algorithm_1.CalculateDiscount(
			testCase.params[0], testCase.params[1],
			testCase.params[2], testCase.params[3],
			testCase.params[4],
		)
		if result != testCase.expected {
			t.Errorf("#%d:\n got: %q\nwant: %q", i, result, testCase.expected)
		}
	}
}
