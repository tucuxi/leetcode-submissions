/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
    rest := limit - root.Val
    if root.Left == nil && root.Right == nil {
        if rest > 0 {
            return nil
        }
        return root
    }
    if root.Left != nil {
        root.Left = sufficientSubset(root.Left, rest)
    }
    if root.Right != nil {
        root.Right = sufficientSubset(root.Right, rest)
    }
    if root.Left == nil && root.Right == nil {
        return nil
    }
    return root
}