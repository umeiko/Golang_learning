package main

import "fmt"

func twoSum(nums []int, target int) []int {
	s := len(nums)
	var out []int
	for i := 0; i < s; i++ {
		for j := i + 1; j < s; j++ {
			if nums[i]+nums[j] == target {
				out = append(out, i, j)
				break
			}
		}
	}
	return out
}

func main() {
	var input []int
	input = append(input, 1, 2, 3)
	fmt.Println(twoSum(input, 4))
}
