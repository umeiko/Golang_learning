package main

import (
	"fmt"
	"golangLearning/tree_node"
)

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func recursion_builder(nums []int, root *tree_node.TreeNode) bool {
	max := 0
	max_index := 0
	if len(nums) == 0 {
		return true
	}
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
			max_index = i
		}
	}
	root.Val = max
	left_nums := nums[:max_index]
	right_nums := nums[max_index+1:]
	if len(left_nums) != 0 {
		left_node := &tree_node.TreeNode{}
		root.Left = left_node
		recursion_builder(left_nums, left_node)
	}
	if len(right_nums) != 0 {
		right_node := &tree_node.TreeNode{}
		root.Right = right_node
		recursion_builder(nums[max_index+1:], right_node)
	}
	return false
}

// 构建最大二叉树
func constructMaximumBinaryTree(nums []int) *tree_node.TreeNode {
	root := &tree_node.TreeNode{}
	recursion_builder(nums, root)
	return root
}

func main() {
	a := []int{3, 2, 1, 6, 0, 5}
	// input := tree_node.GenTest(a)
	result := constructMaximumBinaryTree(a)
	fmt.Print(tree_node.LevelOrderTraversal(result))
}
