package kmp

import (
	"fmt"
	"testing"
)

/*
判断一颗二叉树是是否是另一颗树的子树。比如tree2是tree1的子树。
将过程分为两步：
1，通过递归，在树1中每次给出一个节点。
2，通过递归，在树2中判断该结点所在树是否与树2一样。
*/

// 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestSubtree(t1 *testing.T) {
	// 构造二叉树s
	s := &TreeNode{Val: 3}
	s.Left = &TreeNode{Val: 4}
	s.Right = &TreeNode{Val: 5}
	s.Left.Left = &TreeNode{Val: 1}
	s.Left.Right = &TreeNode{Val: 2}

	// 构造二叉树t
	t := &TreeNode{Val: 4}
	t.Left = &TreeNode{Val: 1}
	t.Right = &TreeNode{Val: 2}

	// 判断t是否是s的子树
	fmt.Println(isSubtree(s, t)) // true
}

// 判断两棵树是否相等
func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

// 判断t是否是s的子树
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if isSameTree(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

/*
方法2：
前序遍历树1，结果集放入数组1
前序遍历树2，结果集放入数组2
数组2如果是数组1的子集，则说明树2是树1的子树。
*/
