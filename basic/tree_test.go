package basic

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历二叉树
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 中序遍历二叉树
func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left)  // 遍历左子树
	fmt.Printf("%d ", root.Val)  // 访问根节点
	inorderTraversal(root.Right) // 遍历右子树
}

// 后序遍历二叉树
func postorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	postorderTraversal(root.Left)
	postorderTraversal(root.Right)
	fmt.Printf("%d ", root.Val)
}

func TestTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	preorderTraversal(root)
	t.Log()
	inorderTraversal(root)
	t.Log()
	postorderTraversal(root)
}
