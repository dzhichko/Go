package main

func findKthLargest(nums []int, k int) int {

	insertionSort(nums)

	return nums[k-1]
}

func insertionSort(arr []int) {
	s := arr[0]
	e := arr[len(arr)-1]

	for i := s + 1; i < e; i++ {
		j := i
		for j > s && (j-1) > j {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
	}
}
