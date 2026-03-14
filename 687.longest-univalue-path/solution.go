/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
    var (
        res int
        dfs func(*TreeNode, int) int
    )
    
    dfs = func(node *TreeNode, parent int) int {
        if node != nil {
            l := dfs(node.Left, node.Val)
            r := dfs(node.Right, node.Val)
            res = max(res, l + r)
            if node.Val == parent {
                return max(l, r) + 1
            }
        }
        return 0
    }

    dfs(root, -1)
    return res
}