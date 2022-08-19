package tree_node

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const Null int = -1

func creat_and_link_node(root *TreeNode, val, L_or_R int) *TreeNode {
	s_root := ""
	s_val := ""
	s_LR := ""
	if root != nil {
		s_root = strconv.Itoa(root.Val)
	} else {
		s_root = "nil"
	}
	if val != Null {
		s_val = strconv.Itoa(val)
	} else {
		s_val = "nil"
	}
	if L_or_R == 0 {
		s_LR = "Left "
	} else {
		s_LR = "Right"
	}
	fmt.Printf("	root:%s\t%s-Node: %s\n", s_root, s_LR, s_val)
	if root != nil && val != Null {
		new_node := &TreeNode{Val: val}
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

func GenTest(val_list []int) *TreeNode {
	num_rest_node := len(val_list) - 1
	root := &TreeNode{}
	if len(val_list) > 0 {
		fmt.Printf("level %d:\n", 0)
		if val_list[0] != Null {
			root.Val = val_list[0]
			fmt.Printf("	Node %d\n", root.Val)
		}

	}
	last_level_Nodes := []*TreeNode{root}
	for i := 1; ; i++ {
		if num_rest_node <= 0 {
			break
		}
		fmt.Printf("level %d:\n", i)
		num_node_in_cur_level := len(last_level_Nodes) * 2
		for num_rest_node < num_node_in_cur_level {
			val_list = append(val_list, Null)
			num_rest_node += 1
		}
		this_level_Nodes := []*TreeNode{}
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
