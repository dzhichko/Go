package main

func findKthLargest(nums []int, k int) int {

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
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] >= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}
