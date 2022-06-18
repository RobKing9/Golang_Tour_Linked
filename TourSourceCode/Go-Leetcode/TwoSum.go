package main

import "fmt"

func TwoSum1( nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; i++{
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(TwoSum1([]int{2, 11, 7, 25}, 9))
}