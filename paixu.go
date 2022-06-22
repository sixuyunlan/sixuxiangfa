package main

import (
	"fmt"
)

func main() {
	fmt.Println(QuickSort([]int{5, 2, 3, 9, 1, 4, 8, 7}))
}

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	low := make([]int, 0, 0)
	high := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	flag := nums[0]
	mid = append(mid, flag)
	for i := 1; i < len(nums); i++ {
		if nums[i] < flag {
			low = append(low, nums[i])

		} else if nums[i] > flag {
			high = append(high, nums[i])
		} else {
			mid = append(mid, nums[i])
		}
	}
	low, high = QuickSort(low), QuickSort(high)
	return append(append(low, mid...), high...)
}
