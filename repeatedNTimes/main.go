package main

import (
	"fmt"
)

func repeatedNTimes(nums []int) (r int) {
	m := make(map[int]int, 5001)
	target := len(nums) / 2
	r = -1
	// for i := 0; i < len(nums); i++ {
	// 	m[nums[i]] = 0
	// }
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k := range m {
		fmt.Print(k)
		if m[k] == target {
			r = k
			return
		}
	}
	return
}

func main() {
	var nums = []int{5, 1, 5, 2, 5, 3, 5, 4}

	fmt.Print(repeatedNTimes(nums))

}
