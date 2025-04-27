/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
    dfs(root.Left, root.Right, true)
    return root
}

func dfs(left, right *TreeNode, oddLevel bool) {
    if left == nil || right == nil {
        return
    }
    if oddLevel {
        left.Val, right.Val = right.Val, left.Val
    }
    dfs(left.Left, right.Right, !oddLevel)
    dfs(left.Right, right.Left, !oddLevel)
}