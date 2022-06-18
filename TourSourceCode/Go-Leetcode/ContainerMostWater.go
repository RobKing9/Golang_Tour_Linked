package main

import "fmt"

func MaxArea1 (height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	maxMulti := 0
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			w := j-1
			h := Min(height[i], height[j])
			maxMulti = Max(maxMulti, w*h)
		}
	}
	return maxMulti
}

func MaxArea2 (height []int) int {
	maxMulti := 0
	left, right := 0, len(height)-1
	for left < right {
		w := right - left
		h := Min(height[left], height[right])
		maxMulti = Max (maxMulti, w*h)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxMulti
}


func Max (vals...int) int {
	var max int
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func Min (vals...int) int {
	var min int
	for _, val := range vals {
		if min == 0||val <= min {
			min = val
		}
	}
	return min
}

func main () {
	fmt.Println(Max(15, 7))
	fmt.Println(Min(15, 7))
	fmt.Println(MaxArea1([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(MaxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}