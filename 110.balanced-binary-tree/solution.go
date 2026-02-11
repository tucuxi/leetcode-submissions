/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    _, b := dfs(root)
    return b
}

func dfs(root *TreeNode) (int, bool) {
    if root == nil {
        return 0, true
    }
    dl, bl := dfs(root.Left)
    dr, br := dfs(root.Right)
    dd := dl - dr
    return max(dl, dr) + 1, bl && br && dd <= 1 && dd >= -1
}