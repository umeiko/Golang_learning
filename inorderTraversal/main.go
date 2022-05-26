package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

}

func main() {
	var nums = []int{5, 1, 5, 2, 5, 3, 5, 4}

	fmt.Print(inorderTraversal(nums))

}
