package main

/**
303. 区域和检索 - 数组不可变
*/
type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	// 求前缀和，下标从0开始，最后求区间[l,r]的和是s[r+1]-s[l]
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}
	return NumArray{sums}
}

func (na *NumArray) SumRange(i, j int) int {
	return na.sums[j+1] - na.sums[i]
}
