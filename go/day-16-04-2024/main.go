package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
			return &TreeNode{Val: val, Left: root}
	}

	var add func(node *TreeNode, currentDepth int)
	add = func(node *TreeNode, currentDepth int) {
			if node == nil {
					return
			}
			if currentDepth == depth-1 {
					newLeft := &TreeNode{Val: val, Left: node.Left}
					newRight := &TreeNode{Val: val, Right: node.Right}
					node.Left = newLeft
					node.Right = newRight
			} else {
					add(node.Left, currentDepth+1)
					add(node.Right, currentDepth+1)
			}
	}

	add(root, 1)
	return root
}

func main() {
	fmt.Println()
}