package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	isLeaf := func(node *TreeNode) bool {
		return node != nil && node.Left == nil && node.Right == nil
	}

	var dfs func(node *TreeNode, isLeft bool) int
	dfs = func(node *TreeNode, isLeft bool) int {
		if node == nil {
				return 0
		}
		if isLeaf(node) && isLeft {
				return node.Val
		}
		return dfs(node.Left, true) + dfs(node.Right, false)
	}

	return dfs(root, false)
}

func main() {
	fmt.Println(sumOfLeftLeaves([]int {3,9,20,null,null,15,7}))
}