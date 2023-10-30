package main

import "fmt"

func findKthLargest(nums []int, k int) int {

	if len(nums) == 0 || k-1 > len(nums) {
		fmt.Println("Invalid length of array")
	}
	nums = mergeSort(nums)
	return nums[k-1]
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(right) == 0 || len(left) > 0 && left[0] >= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}
