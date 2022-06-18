package main

import "fmt"

func merge (nums1, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	i, j := 0, 0
	res := make([]int, 0, n1 + n2)
	for i < n1 && j < n2 {
		switch {
		case nums1[i] < nums2[j]:
			res = append(res, nums1[i])
			i++
		case nums1[i] > nums2[j]:
			res = append(res, nums2[j])
			j++
		default:
			res = append(res, nums1[i], nums2[j])
			i++
			j++
		}
	}

	if i < n1 {
		res = append(res, nums1[i:]...)
	}
	if j < n2 {
		res = append(res, nums2[j:]...)
	}
	return res
}

func MedianOfArrays (nums1 []int, nums2 []int) float64 {
	res := merge(nums1, nums2)
	fmt.Println(res)
	n := len(res)
	if n == 0 {
		return -1
	}
	if n%2 == 0{
		return float64(res[n/2-1]+res[n/2]) / 2
	}
	return float64(res[n/2])
}

func main () {
	a := MedianOfArrays([]int{1, 3, 5}, []int{2, 4, 6})
	fmt.Println(a)
}
