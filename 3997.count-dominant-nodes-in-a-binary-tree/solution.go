/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countDominantNodes(root *TreeNode) int {
    res := 0

    var dfs func(*TreeNode) int

    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        m := max(dfs(node.Left), dfs(node.Right), node.Val)
        if node.Val == m {
            res++
        }
        return m
    }

    dfs(root)
    return res
}