package main

import (
	"fmt"
	"inorderTraversal/tree_node"
)

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

var null int = -1

func creat_and_link_node(root *tree_node.TreeNode, val, L_or_R int) *tree_node.TreeNode {
	if root != nil {
		fmt.Printf("	root:%d, %d-Node:%d\n", root.Val, L_or_R, val)
	} else {
		fmt.Printf("	root:nil, %d-Node:%d\n", L_or_R, val)
	}
	if root != nil && val != null {
		new_node := &tree_node.TreeNode{Val: val}
		if L_or_R == 0 {
			root.Left = new_node
		} else if L_or_R == 1 {
			root.Right = new_node
		}
		return new_node
	} else {
		return nil
	}
}

func genTest(val_list []int) *tree_node.TreeNode {
	num_rest_node := len(val_list) - 1
	root := &tree_node.TreeNode{}
	if len(val_list) > 0 {
		if val_list[0] != null {
			root.Val = val_list[0]
		}

	}
	last_level_Nodes := []*tree_node.TreeNode{root}
	for i := 1; ; i++ {
		if num_rest_node <= 0 {
			break
		}
		fmt.Printf("level %d:\n", i)
		num_node_in_cur_level := len(last_level_Nodes) * 2
		for num_rest_node < num_node_in_cur_level {
			val_list = append(val_list, null)
			num_rest_node += 1
		}
		this_level_Nodes := []*tree_node.TreeNode{}
		for j := 0; j < num_node_in_cur_level; j++ {
			new_node := creat_and_link_node(last_level_Nodes[j/2], val_list[len(val_list)-num_rest_node], j%2)
			if new_node != nil {
				this_level_Nodes = append(this_level_Nodes, new_node)
			}
			num_rest_node -= 1
		}
		last_level_Nodes = this_level_Nodes
	}
	return root
}

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
	a := []int{1, 2, 3, 4, 5, 6, 7}
	root := genTest(a)
	result := deepestLeavesSum(root)
	fmt.Print(result)
	// a := 1
	// fmt.Print(a / 2)
}
