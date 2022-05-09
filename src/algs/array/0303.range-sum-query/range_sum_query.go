package main

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}
	return NumArray{sums}
}

func (na *NumArray) SumRange(left int, right int) int {
	return na.sums[right+1] - na.sums[left]
}
