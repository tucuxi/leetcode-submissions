/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt

    var walk func(*TreeNode) int
    walk = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        l := max(0, walk(root.Left))
        r := max(0, walk(root.Right))
        maxSum = max(maxSum, root.Val + l + r)
        return root.Val + max(l, r)
    }

    walk(root)
    return maxSum
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}