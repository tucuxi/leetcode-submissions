/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    return slices.Concat(postorderTraversal(root.Left), postorderTraversal(root.Right), []int{root.Val}) 
}