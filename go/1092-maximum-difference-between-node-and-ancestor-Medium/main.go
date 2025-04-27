/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
    diff := 0

    var walk func(*TreeNode, int, int)
    walk = func(root *TreeNode, min, max int) {
        if root == nil {
            return
        }
        if root.Val < min {
            min = root.Val
        }
        if root.Val > max {
            max = root.Val
        }
        if d := max - min; d > diff {
            diff = d
        }
        walk(root.Left, min, max)
        walk(root.Right, min, max)
    }

    walk(root, math.MaxInt, math.MinInt)
    return diff
}

