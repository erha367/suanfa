package suanfa

import (
	"testing"
)

//给定一个链表，如果有环路，找出环路的开始点。

func TestLinkedTestCycle(t *testing.T) {
	p1 := &Node{1, nil}
	p2 := &Node{2, nil}
	p3 := &Node{3, nil}
	p4 := &Node{4, nil}
	p5 := &Node{5, nil}
	p1.Next = p2
	p2.Next = p3
	p3.Next = p4
	p4.Next = p5
	p5.Next = p3
	t.Log(hasCycle(p1))
}

type Node struct {
	Val  int
	Next *Node
}

func hasCycle(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	var cycle bool
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			cycle = true
			break
		}
	}
	if cycle == false {
		return nil
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
