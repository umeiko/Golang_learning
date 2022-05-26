package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将数组类型转为链表，返回链表指针
func makeListNone(nums []int) (out *ListNode) {
	// 创建结构体并申请内存空间
	preNode := &ListNode{Next: nil}
	for i := range nums {
		newNode := &ListNode{Val: 0}
		// 如果在没有申请空间的情况下赋值会导致panic错误
		newNode.Val = nums[i]
		preNode.Next = newNode
		preNode = newNode
		if i == 0 {
			out = newNode
		}
	}
	return
}

// 将链表转换为数组
func decodeListNode(in *ListNode) (out []int) {
	for {
		if in == nil {
			break
		}
		out = append(out, in.Val)
		in = in.Next
	}
	return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (out *ListNode) {
	sum_0, sum_1 := 0, 0
	head := true
	preNode := &ListNode{Next: nil}
	for {
		num1, n1 := getValueAndNextNode(l1)
		num2, n2 := getValueAndNextNode(l2)
		l1 = n1
		l2 = n2
		temp := (num1 + num2 + sum_1)
		sum_0 = temp % 10
		sum_1 = temp / 10

		nowNode := &ListNode{Val: sum_0}
		preNode.Next = nowNode
		preNode = nowNode
		if head {
			out = nowNode
		}
		head = false
		if l1 == nil && l2 == nil && sum_1 == 0 {
			break
		}
	}
	return

}

// 取出链表中的值，并返回下一个链表指针，如果是空链表则值为0
func getValueAndNextNode(in *ListNode) (v int, outNode *ListNode) {
	if in != nil {
		v = in.Val
		outNode = in.Next
	} else {
		v = 0
	}
	return
}

func main() {
	var nums1 = makeListNone([]int{2, 4, 3})
	var nums2 = makeListNone([]int{5, 6, 6})

	// var blankNode ListNode
	result := addTwoNumbers(nums1, nums2)
	fmt.Println(decodeListNode(result))

}
