package main

func findKthLargest(nums []int, k int) int {
	insertionSort(nums)
	return nums[k-1]
}
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}
