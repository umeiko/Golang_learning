package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func genTest() *TreeNode {
	root := &TreeNode{Val: 1}
	R := &TreeNode{Val: 2}
	RL := &TreeNode{Val: 3}
	RR := &TreeNode{Val: 10}
	L := &TreeNode{Val: 3}
	LL := &TreeNode{Val: 5}
	LR := &TreeNode{Val: 7}
	root.Right = R
	root.Left = L
	R.Left = RL
	R.Right = RR
	L.Left = LL
	L.Right = LR
	return root
}

func genTest2([]int) *TreeNode {
	i := 0
	root := &TreeNode{}
	for i < 5 {
		fmt.Println(i)
		for j := 0; j < int(math.Pow(2., float64(i))); j++ {
			fmt.Print(j)
		}
		fmt.Println()
		i++
	}
	return root
}

func inorderTraversal(root *TreeNode) (out []int) {
	// now_node := root
	// pre_node := root
	// if pre_node.Left != nil {
	// 	now_node = pre_node.Left
	// } else {
	// 	now_node = nil
	// }
	if root == nil {
		return
	}
	if (root.Left == nil) && (root.Right == nil) {
		out = append(out, root.Val)
		return
	}
	for {
		pre_node := root
		// fmt.Print(pre_node.Left)
		// fmt.Println(pre_node.Right)
		for {
			// fmt.Print(pre_node.Val)
			if pre_node.Left != nil {
				now_node := pre_node.Left
				if now_node.Left != nil {
					pre_node = now_node
				} else {
					out = append(out, now_node.Val)
					pre_node.Left = now_node.Right
					if root.Val != 101 && (root.Left == nil) && (root.Right == nil) {
						out = append(out, root.Val)
						return
					}
					break
				}
			} else {
				out = append(out, pre_node.Val)
				pre_node.Val = 101
				pre_node.Left = pre_node.Right
				pre_node.Right = nil
			}
		}
		// fmt.Print(root.Left)
		// fmt.Println(root.Right)
		// fmt.Println()
		if (root.Left == nil) && (root.Right == nil) {
			// fmt.Print("Done?")
			return
		}
	}
	// return
}

func main() {
	// a := []int{}
	// genTest2(a)
	fmt.Print(inorderTraversal(genTest()))

}
