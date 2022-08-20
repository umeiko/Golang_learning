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

func get_node(root *TreeNode, L_or_R int) (out *TreeNode) {
	if L_or_R == 0 {
		out = root.Left
	}
	if L_or_R == 1 {
		out = root.Right
	}
	return
}

// 输入数组，自动生成测试数列
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

// 中序遍历
func InorderTraversal(root *TreeNode) (out []int) {
	if root == nil {
		return
	}
	if (root.Left == nil) && (root.Right == nil) {
		out = append(out, root.Val)
		return
	}
	for {
		pre_node := root
		for {
			if pre_node.Left != nil {
				now_node := pre_node.Left
				if now_node.Left != nil {
					pre_node = now_node
				} else {
					out = append(out, now_node.Val)
					pre_node.Left = now_node.Right
					if root.Val != Null && (root.Left == nil) && (root.Right == nil) {
						out = append(out, root.Val)
						return
					}
					break
				}
			} else {
				out = append(out, pre_node.Val)
				pre_node.Val = Null
				pre_node.Left = pre_node.Right
				pre_node.Right = nil
			}
		}

		if (root.Left == nil) && (root.Right == nil) {
			return
		}
	}
}

// 层序遍历获取二叉树
func LevelOrderTraversal(root *TreeNode) [][]int {
	now_nodes := []*TreeNode{}
	new_values := []int{}
	out := [][]int{}
	if root != nil {
		now_nodes = append(now_nodes, root)
		new_values = append(new_values, root.Val)
		out = append(out, new_values)
	}
	for len(now_nodes) != 0 {
		new_nodes := []*TreeNode{}
		new_values = []int{}
		for i := 0; i < 2*len(now_nodes); i++ {
			next_node := get_node(now_nodes[i/2], i%2)
			if next_node != nil {
				new_values = append(new_values, next_node.Val)
				new_nodes = append(new_nodes, next_node)
			}
		}
		if len(new_values) != 0 {
			out = append(out, new_values)
		}
		now_nodes = new_nodes
	}
	return out
}
