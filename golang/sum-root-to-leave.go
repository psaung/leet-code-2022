package main

import "fmt"
import "strconv"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}





func sumRootToLeaf(root *TreeNode) int {
	result := 0
    

	var bfs func (*TreeNode, int)

	bfs = func(node *TreeNode, a int) {
		if node == nil {
			return
		}

        a = a << 1 + node.Val
		if node.Left == nil && node.Right == nil {
            
            result += a
            return
		} 
        
		bfs(node.Left, a)

		bfs(node.Right, a)
	}

    bfs(root, 0)

	return result;
}