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

const null int = -1

func get_all_next_node(nodes []*tree_node.TreeNode) (out []*tree_node.TreeNode) {
	for i := range nodes {
		if nodes[i] != nil {
			if nodes[i].Left != nil {
				out = append(out, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				out = append(out, nodes[i].Right)
			}
		}
	}
	return out
}

func deepestLeavesSum(root *tree_node.TreeNode) (out int) {
	b := []*tree_node.TreeNode{root}
	temp := b
	for {
		temp = get_all_next_node(b)
		if len(temp) != 0 {
			b = temp
		} else {
			break
		}
	}
	for i := range b {
		out += b[i].Val
	}
	return
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, null}
	root := tree_node.GenTest(a)
	result := deepestLeavesSum(root)
	fmt.Print(result)
	// a := 1
	// fmt.Print(a / 2)
}
