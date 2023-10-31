package tree

import (
	"strconv"
	"strings"
	"testing"
)

/*
题目：给你一个二叉树，序列化从一个string字符串，再从一个string字符串还原出二叉树
1.序列化：前序遍历，根左右
2.反序列化：递归，根左右
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestSerialize(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	res := serialize(root)
	t.Log(res)

	res2 := deserialize(res)
	t.Log(res2.Val)
}

func serialize(root *TreeNode) string {
	//前序遍历（根->左->右）
	if root == nil {
		return "# "
	}
	res := strconv.Itoa(root.Val) + " "
	//左节点
	res += serialize(root.Left)
	//右节点
	res += serialize(root.Right)
	return res
}

/*-- 反序列化 --*/

func deserialize(data string) *TreeNode {
	nodes := strings.Split(data, " ")
	return deserializeHelper(&nodes)
}

func deserializeHelper(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}
	if (*nodes)[0] == "#" {
		*nodes = (*nodes)[1:]
		return nil
	}
	val, _ := strconv.Atoi((*nodes)[0])
	*nodes = (*nodes)[1:]
	return &TreeNode{
		Val:   val,
		Left:  deserializeHelper(nodes),
		Right: deserializeHelper(nodes),
	}
}
