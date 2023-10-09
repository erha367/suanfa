package suanfa

import (
	"fmt"
	"testing"
)

func TestNodes(t *testing.T) {
	p3 := &Node{3, nil}
	p2 := &Node{2, p3}
	p1 := &Node{1, p2}
	PrintNode(p1)
	t.Log(`--------------------------------`)
	PrintNode(ReverseNodes(p1))
}

// 反转链表
func ReverseNodes(n *Node) *Node {
	var prev *Node
	for n != nil {
		next := n.Next
		n.Next = prev
		prev = n
		n = next
	}
	return prev
}

func PrintNode(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.Val)
	PrintNode(n.Next)
}

type Node struct {
	Val  int
	Next *Node
}
