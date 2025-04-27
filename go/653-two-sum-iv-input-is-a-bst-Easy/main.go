/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    vals := map[int]bool{}

    var walk func(*TreeNode) bool
    walk = func(root *TreeNode) bool {
        if root == nil {
            return false
        }
        if vals[k - root.Val] {
            return true
        }
        vals[root.Val] = true
        return walk(root.Left) || walk(root.Right)
    }

    return walk(root)
}