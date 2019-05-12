package main

import "fmt"

/**
插入排序
*/
func insert(nums []int) []int {
	for j := 1; j < len(nums); j++ {
		key := nums[j]
		i := j - 1
		for i >= 0 && nums[i] > key {
			nums[i+1] = nums[i]
			i -= 1
		}
		nums[i+1] = key
	}
	return nums
}

func main() {
	ret := insert([]int{4, 5, 3, 7, 1, 10})
	fmt.Print(ret)
}
