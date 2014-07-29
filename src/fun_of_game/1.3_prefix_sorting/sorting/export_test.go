// Bridge package to expose sorting internals to tests
package sorting

type InitResult struct {
	CakeArray            []int
	CakeCnt              int
	MaxSwap              int
	SwapArray            []int
	ReverseCakeArray     []int
	ReverseCakeArraySwap []int
}

func (sorting PrefixSorting) Init(pCakeArray []int, nCakeCnt int) InitResult {

	sorting.init(pCakeArray, nCakeCnt)

	return InitResult{
		sorting.m_CakeArray,
		sorting.m_nCakeCnt,
		sorting.m_nMaxSwap,
		sorting.m_SwapArray,
		sorting.m_ReverseCakeArray,
		sorting.m_ReverseCakeArraySwap,
	}
}

func (sorting PrefixSorting) UpBound(nCakeCnt int) int {
	return sorting.upBound(nCakeCnt)
}

func (sorting PrefixSorting) LowerBound(pCakeArray []int, nCakeCnt int) int {
	return sorting.lowerBound(pCakeArray, nCakeCnt)
}

func (sorting PrefixSorting) IsSorted(pCakeArray []int, nCakeCnt int) bool {
	return sorting.isSorted(pCakeArray, nCakeCnt)
}

func (sorting PrefixSorting) Revert(pCakeArray []int, nCakeCnt int, nBegin int, nEnd int) []int {
	sorting.m_ReverseCakeArray = make([]int, nCakeCnt, nCakeCnt)
	for i := 0; i < nCakeCnt; i++ {
		sorting.m_ReverseCakeArray[i] = pCakeArray[i]
	}
	sorting.revert(nBegin, nEnd)

	return sorting.m_ReverseCakeArray
}
