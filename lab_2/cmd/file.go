package main

func findKthLargest(nums []int, k int) int {
	nums = quickSort(nums)
	return nums[len(nums)-k]
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	start := arr[0]
	var low, top []int
	for _, num := range arr[1:] {
		if num <= start {
			low = append(low, num)
		} else {
			top = append(top, num)
		}
	}
	res := append(quickSort(low), start)
	res = append(res, quickSort(top)...)
	return res
}
