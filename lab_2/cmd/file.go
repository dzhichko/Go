package main

func findKthLargest(nums []int, k int) int {
	nums = quickSort(nums)
	return nums[k-1]
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	start := arr[0]
	var low, top []int
	for _, num := range arr[1:] {
		if num >= start {
			top = append(top, num)
		} else {
			low = append(low, num)
		}
	}
	res := append(quickSort(top), start)
	res = append(res, quickSort(low)...)
	return res
}
