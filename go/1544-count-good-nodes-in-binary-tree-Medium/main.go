/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    count := 0
    walk(root, math.MinInt, func() { count++ })
    return count
}

func walk(node *TreeNode, max int, action func()) {
    if node == nil {
        return
    }
    if node.Val >= max {
        max = node.Val
        action()
    }
    walk(node.Left, max, action)
    walk(node.Right, max, action)
}