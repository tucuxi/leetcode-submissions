/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0

    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        l1 := dfs(root.Left)
        l2 := dfs(root.Right)
        diameter = max(diameter, l1 + l2)
        return max(l1, l2) + 1
    }

    dfs(root)
    return diameter
}