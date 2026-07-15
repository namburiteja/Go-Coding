package main

import "fmt"

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	nums := []int{2, 5, 8, 10, 14, 18, 20}
	target := 14

	index := binarySearch(nums, target)

	if index != -1 {
		fmt.Println("Element found at index:", index)
	} else {
		fmt.Println("Element not found")
	}
}