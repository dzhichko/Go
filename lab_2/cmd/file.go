package main

func findKthLargest(nums []int, k int) int {

	insertionSort(nums)

	return nums[k-1]
}

func insertionSort(arr []int) {
	e := len(arr)

	for i := 1; i < e; i++ {
		for j := 0; j < i; j++ {
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
}
