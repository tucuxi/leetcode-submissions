/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
    var (
        ans int
        dfs func(*TreeNode) int
    )
    dfs = func(p *TreeNode) int {
        if p == nil {
            return 0
        }
        l := dfs(p.Left)
        r := dfs(p.Right)
        ans += abs(l) + abs(r)
        return p.Val - 1 + l + r
    }
    dfs(root)
    return ans
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}