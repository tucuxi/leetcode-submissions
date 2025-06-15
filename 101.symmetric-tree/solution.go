/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return root == nil || eq(root.Left, root.Right)
}

func eq(root1, root2 *TreeNode) bool {
    if root1 == nil || root2 == nil {
        return root1 == root2
    }
    return root1.Val == root2.Val &&
        eq(root1.Left, root2.Right) &&
        eq(root1.Right, root2.Left)
}