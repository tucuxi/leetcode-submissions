/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    res := 0
    counter := 0

    var inorder func(*TreeNode)
    inorder = func(root *TreeNode) {
        if root == nil || counter >= k {
            return
        }
        inorder(root.Left)
        if counter++; counter == k {
            res = root.Val
        }
        inorder(root.Right)        
    }

    inorder(root)
    return res
}