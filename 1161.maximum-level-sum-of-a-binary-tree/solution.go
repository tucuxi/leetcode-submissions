/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    maxSum := math.MinInt
    maxLevel := 0
    level := 0
    q := []*TreeNode{root}
    for len(q) > 0 {
        level++
        levelSum := 0
        for i := len(q); i > 0; i-- {
            node := q[0]
            q = q[1:]
            levelSum += node.Val
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        if levelSum > maxSum {
            maxSum = levelSum
            maxLevel = level
        }
    }
    return maxLevel
}