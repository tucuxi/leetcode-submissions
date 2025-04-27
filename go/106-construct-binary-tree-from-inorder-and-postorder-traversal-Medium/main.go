/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    n := len(postorder)
    if n == 0 {
        return nil
    }
    v := postorder[n - 1]
    k := 0
    for inorder[k] != v {
        k++
    }
    return &TreeNode{v,
        buildTree(inorder[:k], postorder[:k]),
        buildTree(inorder[k+1:], postorder[k:n-1])}
}