package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, exists := hashMap[complement]; exists {
			return []int{index, i}
		}

		hashMap[num] = i
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	fmt.Println("Indices:", result)
}