/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    return sumLeft(root, false)
}

func sumLeft(node *TreeNode, left bool) int {
    if node == nil {
        return 0
    }
    if node.Left == nil && node.Right == nil && left {
        return node.Val
    }
    return sumLeft(node.Left, true) + sumLeft(node.Right, false)
}