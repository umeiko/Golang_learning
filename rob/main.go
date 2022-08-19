package main

import (
	"fmt"
)

func rob(nums []int) (r int) {
	// 第n间房间的最大收益 = max(这间 + 上上间的最大收益, 上间的最大收益)
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	max_value := make([]int, len(nums))
	max_value[0] = nums[0]
	max_value[1] = max(nums[0], nums[1])

	for i := 2; i < len(max_value); i++ {
		max_value[i] = max(nums[i]+max_value[i-2], max_value[i-1])
	}
	return max_value[len(nums)-1]
}

func max(x, y int) (r int) {
	if x > y {
		r = x
	} else {
		r = y
	}
	return
}

func main() {
	var nums = []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))

}
