/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    levels := [][]int{}
    
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, level int) {
        if node == nil {
            return
        }
        if len(levels) == level {
            levels = append(levels, []int{node.Val})
        } else if level % 2 == 0 {
            levels[level] = append(levels[level], node.Val)
        } else {
            levels[level] = append([]int{node.Val}, levels[level]...)
        }
        dfs(node.Left, level + 1)
        dfs(node.Right, level + 1)
    }

    dfs(root, 0)
    return levels
}