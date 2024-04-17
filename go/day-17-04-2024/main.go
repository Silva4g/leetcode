package main

import "fmt"

func smallestFromLeaf(root *TreeNode) string {
	smallest := "~"

	var dfs func(node *TreeNode, path []byte)
	dfs = func(node *TreeNode, path []byte) {
			if node == nil {
					return
			}
			path = append([]byte{byte(node.Val + 'a')}, path...)
			if node.Left == nil && node.Right == nil {
					candidate := string(path)
					if candidate < smallest {
							smallest = candidate
					}
			}
			dfs(node.Left, path)
			dfs(node.Right, path)
	}
	dfs(root, []byte{})

	return smallest
}

func main() {
	fmt.Println()
}