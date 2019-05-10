package main

import (
	"math/rand"
	"time"
)

func shuffle(nums []int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(nums); i > 0; i-- {
		last := i - 1
		idx := rand.Intn(i)
		nums[last], nums[idx] = nums[idx], nums[last]
	}
	return nums
}
