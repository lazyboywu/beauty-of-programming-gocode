package sorting_test

import (
	//"fmt"
	. "fun_of_game/1.3_prefix_sorting/sorting"
	"reflect"
	"testing"
)

var upBoundTests = []struct {
	input    int
	expected int
}{
	{1, 2},
	{0, 0},
}

func TestUpBound(t *testing.T) {
	s := NewPrefixSorting()
	//defer s.Close()

	for i, test := range upBoundTests {
		result := s.UpBound(test.input)
		if result != test.expected {
			t.Errorf("#%d:\n got: %q\nwant: %q", i, result, test.expected)
		}
	}
}

var lowerBoundTests = []struct {
	array    []int
	n        int
	expected int
}{
	{[]int{2, 1, 3}, 3, 1},
	{[]int{1, 2, 3}, 3, 0},

	{[]int{1, 4, 2, 3}, 4, 2},
}

func TestLowerBound(t *testing.T) {
	s := NewPrefixSorting()

	for i, test := range lowerBoundTests {
		result := s.LowerBound(test.array, test.n)
		if result != test.expected {
			t.Errorf("#%d:\n got: %q\nwant: %q", i, result, test.expected)
		}
	}
}

var isSortedTests = []struct {
	array    []int
	n        int
	expected bool
}{
	{[]int{2, 1, 3}, 3, false},
	{[]int{1, 2, 3}, 3, true},

	{[]int{1, 4, 2, 3}, 3, false},
}

func TestIsSorted(t *testing.T) {
	s := NewPrefixSorting()

	for i, test := range isSortedTests {
		result := s.IsSorted(test.array, test.n)
		if result != test.expected {
			t.Errorf("#%d:\n got: %t\nwant: %t", i, result, test.expected)
		}
	}
}

var revertTests = []struct {
	array    []int
	n        int
	begin    int
	end      int
	expected []int
}{
	{[]int{1, 2, 3}, 3, 0, 1, []int{2, 1, 3}},
	{[]int{1, 2, 3, 4}, 4, 0, 2, []int{3, 2, 1, 4}},
	{[]int{1, 2, 3, 4, 5, 6, 7}, 7, 0, 5, []int{6, 5, 4, 3, 2, 1, 7}},
}

func TestRevert(t *testing.T) {
	s := NewPrefixSorting()

	for i, test := range revertTests {
		result := s.Revert(test.array, test.n, test.begin, test.end)
		if !IntArrayEquals(result, test.expected) {
			t.Errorf("#%d:\n got: %#v\nwant: %#v", i, result, test.expected)
		}
	}
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

var initTests = []struct {
	array    []int
	n        int
	expected InitResult
}{
	{
		[]int{1, 2, 3}, 3, InitResult{
			[]int{1, 2, 3}, 3, 6,
			[]int{0, 0, 0, 0, 0, 0, 0},
			[]int{1, 2, 3}, []int{0, 0, 0, 0, 0, 0},
		},
	},
	{
		[]int{1, 2, 3, 4}, 4, InitResult{
			[]int{1, 2, 3, 4}, 4, 8,
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			[]int{1, 2, 3, 4}, []int{0, 0, 0, 0, 0, 0, 0, 0},
		},
	},
}

func TestInit(t *testing.T) {
	s := NewPrefixSorting()

	for i, test := range initTests {
		result := s.Init(test.array, test.n)
		if !InitResultEquals(result, test.expected) {
			t.Errorf("#%d:\n got: %#v\nwant: %#v", i, result, test.expected)
		}
		s.Close()
	}
}

func InitResultEquals(a InitResult, b InitResult) bool {
	// 简单的方法
	// if reflect.DeepEqual(a, b) {
	// 	return true
	// }

	ra := reflect.ValueOf(a)
	rb := reflect.ValueOf(b)

	for _, name := range [...]string{"CakeArray", "SwapArray", "ReverseCakeArray", "ReverseCakeArraySwap"} {
		if !IntArrayEquals(ra.FieldByName(name).Interface().([]int), rb.FieldByName(name).Interface().([]int)) {
			return false
		}
	}

	for _, name := range [...]string{"CakeCnt", "MaxSwap"} {
		if ra.FieldByName(name).Int() != rb.FieldByName(name).Int() {
			return false
		}
	}

	return true
}
