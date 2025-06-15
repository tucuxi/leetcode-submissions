/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
    if root == nil || val > root.Val {
        return &TreeNode{val, root, nil}
    }
    root.Right = insertIntoMaxTree(root.Right, val)
    return root
}