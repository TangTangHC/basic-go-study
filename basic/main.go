package main

import (
	"fmt"
	"gorm.io/gorm"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) *TreeNode {
	return helper(root, o1, o2)
}

func helper(root *TreeNode, o1 int, o2 int) *TreeNode {
	if root == nil || root.Val == o1 || root.Val == o2 {
		return root
	}
	left := helper(root.Left, o1, o2)
	right := helper(root.Right, o1, o2)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

func main() {
	//root := &TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val: 5,
	//		Left: &TreeNode{
	//			Val:   6,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   2,
	//			Left:  &TreeNode{Val: 7},
	//			Right: &TreeNode{Val: 4},
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:   1,
	//		Left:  &TreeNode{Val: 0},
	//		Right: &TreeNode{Val: 8},
	//	},
	//}
	//
	//// 测试用例：查找节点值为6和4的最低公共祖先节点
	//lca := lowestCommonAncestor(root, 6, 4)
	//fmt.Println("最低公共祖先节点值:", lca.Val)
	var err error
	fmt.Println(err == ErrUserNotFound)
	fmt.Println(err != ErrUserNotFound)
	fmt.Println(err == nil)
}

var ErrUserNotFound = gorm.ErrRecordNotFound
