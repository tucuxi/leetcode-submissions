/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var (
        res []int
        traverse func(*TreeNode)
    )
    
    traverse = func(root *TreeNode) {
        if root != nil {
            traverse(root.Left)
            res = append(res, root.Val)
            traverse(root.Right)
        }
    }
    
    traverse(root)
    return res    
}