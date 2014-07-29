package sorting

import (
	"fmt"
)

type PrefixSorting struct {
	m_CakeArray []int // 烙饼信息数组
	m_nCakeCnt  int   // 烙饼个数
	m_nMaxSwap  int   // 最多交换次数。根据前面的推断，这里最多为 m_nCakeCnt * 2
	m_SwapArray []int // 交换结果数组

	m_ReverseCakeArray     []int // 当前翻转烙饼信息数组
	m_ReverseCakeArraySwap []int // 当前翻转烙饼交换结果数组
	m_nSearch              int   // 当前搜索次数
}

func NewPrefixSorting() *PrefixSorting {
	return &PrefixSorting{m_nCakeCnt: 0, m_nMaxSwap: 0}
}

// @todo 析构函数
func (sorting PrefixSorting) Close() {
	if sorting.m_CakeArray != nil {
		sorting.m_CakeArray = nil
	}
	if sorting.m_SwapArray != nil {
		sorting.m_SwapArray = nil
	}
	if sorting.m_ReverseCakeArray != nil {
		sorting.m_ReverseCakeArray = nil
	}
	if sorting.m_ReverseCakeArraySwap != nil {
		sorting.m_ReverseCakeArraySwap = nil
	}
}

func (sorting *PrefixSorting) Run(pCakeArray []int, nCakeCnt int) {
	sorting.init(pCakeArray, nCakeCnt)

	sorting.m_nSearch = 0
	sorting.search(0)
}

func (sorting *PrefixSorting) Output() {
	for i := 0; i < sorting.m_nMaxSwap; i++ {
		fmt.Printf("%d ", sorting.m_SwapArray[i])
	}
	fmt.Printf("\n |Search Times| : %d\n", sorting.m_nSearch)
	fmt.Printf("Total Swap times = %d\n", sorting.m_nMaxSwap)
}

// 初始化数组信息
// @param
// pCakeArray 存储烙饼索引数组
// nCakeCnt 烙饼个数
func (sorting *PrefixSorting) init(pCakeArray []int, nCakeCnt int) {
	// @todo Assert
	// Assert(pCakeArray != nil)
	// Assert(nCakeCnt > 0)

	sorting.m_nCakeCnt = nCakeCnt

	// 初始化烙饼数组
	sorting.m_CakeArray = make([]int, sorting.m_nCakeCnt)
	// Assert(sorting.m_CakeArray != nil)
	for i := 0; i < sorting.m_nCakeCnt; i++ {
		sorting.m_CakeArray[i] = pCakeArray[i]
	}

	// 设置最多交换次数信息
	sorting.m_nMaxSwap = sorting.upBound(sorting.m_nCakeCnt)

	// 初始化交换结果数组
	sorting.m_SwapArray = make([]int, sorting.m_nMaxSwap+1)
	// Assert(sorting.m_SwapArray != nil)

	// 初始化中间交换结果信息
	sorting.m_ReverseCakeArray = make([]int, sorting.m_nCakeCnt, sorting.m_nCakeCnt)
	for i := 0; i < sorting.m_nCakeCnt; i++ {
		sorting.m_ReverseCakeArray[i] = sorting.m_CakeArray[i]
	}
	sorting.m_ReverseCakeArraySwap = make([]int, sorting.m_nMaxSwap+1)
}

// 寻找当前翻转的上界
func (sorting PrefixSorting) upBound(nCakeCnt int) int {
	return nCakeCnt * 2
}

// 寻找当前翻转的上界
func (sorting PrefixSorting) lowerBound(pCakeArray []int, nCakeCnt int) int {
	t, ret := 0, 0

	// 根据当前数组的排序信息情况来判断最少需要交换多少次
	for i := 1; i < nCakeCnt; i++ {
		// 判断位置相邻的两个烙饼，是否为尺寸排序上相邻的
		t = pCakeArray[i] - pCakeArray[i-1]
		if (t == 1) || (t == -1) {
			// nothing
		} else {
			ret++
		}
	}
	return ret
}

// 排序的主函数
func (sorting *PrefixSorting) search(step int) {
	var nEstimate int

	sorting.m_nSearch++

	// 估算这次搜索所需要的最小交换次数
	nEstimate = sorting.lowerBound(sorting.m_ReverseCakeArray, sorting.m_nCakeCnt)
	if step+nEstimate > sorting.m_nMaxSwap {
		return
	}

	// 如果已经排好序，即翻转完成
	if sorting.isSorted(sorting.m_ReverseCakeArray, sorting.m_nCakeCnt) {
		if step < sorting.m_nMaxSwap {
			sorting.m_nMaxSwap = step
			for i := 0; i < sorting.m_nMaxSwap; i++ {
				sorting.m_SwapArray[i] = sorting.m_ReverseCakeArraySwap[i]
			}
		}
		return
	}

	// 递归进行翻转
	for i := 1; i < sorting.m_nCakeCnt; i++ {
		sorting.revert(0, i)
		sorting.m_ReverseCakeArraySwap[step] = i
		sorting.search(step + 1)
		sorting.revert(0, i)
	}

}

func (sorting PrefixSorting) isSorted(pCakeArray []int, nCakeCnt int) bool {
	for i := 1; i < nCakeCnt; i++ {
		if pCakeArray[i-1] > pCakeArray[i] {
			return false
		}
	}
	return true
}

// 翻转烙饼信息
func (sorting *PrefixSorting) revert(nBegin int, nEnd int) {
	// Assert(nEnd > nBegin)
	var i, j int
	//var t  int

	for i, j = nBegin, nEnd; i < j; i, j = i+1, j-1 {
		//t = sorting.m_ReverseCakeArray[i]
		//sorting.m_ReverseCakeArray[i] = sorting.m_ReverseCakeArray[j]
		//sorting.m_ReverseCakeArray[j] = t

		// so easy
		sorting.m_ReverseCakeArray[i], sorting.m_ReverseCakeArray[j] = sorting.m_ReverseCakeArray[j], sorting.m_ReverseCakeArray[i]
	}
}
