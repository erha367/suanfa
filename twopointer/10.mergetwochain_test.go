package twopointer

import (
	"fmt"
	"testing"
)

// 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建链表
func createList(arr []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, val := range arr {
		node := &ListNode{Val: val}
		current.Next = node
		current = current.Next
	}
	return dummy.Next
}

func TestMergeChain(t *testing.T) {
	l1 := createList([]int{1, 2, 4})
	l2 := createList([]int{1, 3, 4})
	printList(l1)
	printList(l2)
	// 合并两个有序链表
	mergedList := mergeTwoLists(l1, l2)
	// 打印合并后的链表
	printList(mergedList)
}

// 打印链表
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

// 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建一个哑节点作为合并后链表的头节点
	dummy := &ListNode{}
	current := dummy

	// 遍历两个链表，比较节点值，将较小值的节点添加到合并后链表中
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// 将剩余的节点直接添加到合并后链表的末尾
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}
