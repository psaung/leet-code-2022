package main


func inorderTraversal(root *TreeNode) []int {
    
    if root == nil {
        return []int{}
    }
    
    return append(append(append([]int{}, inorderTraversal(root.Left)...) , root.Val), inorderTraversal(root.Right)...)
    
}