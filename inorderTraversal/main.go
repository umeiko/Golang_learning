package inorder_traversal

import (
	"golangLearning/tree_node"
)

const null = tree_node.Null

func InorderTraversal(root *tree_node.TreeNode) (out []int) {
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
					if root.Val != tree_node.Null && (root.Left == nil) && (root.Right == nil) {
						out = append(out, root.Val)
						return
					}
					break
				}
			} else {
				out = append(out, pre_node.Val)
				pre_node.Val = tree_node.Null
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

// func main() {
// 	a := []int{1, 2, 3, null, 4, 5, null, null}
// 	b := tree_node.GenTest(a)
// 	fmt.Print(InorderTraversal(b))

// }
