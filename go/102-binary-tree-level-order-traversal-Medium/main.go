/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    res := make([][]int, 0)
    queue := []*TreeNode{ root }
    for len(queue) > 0 {
        n := len(queue)
        level := make([]int, n)
        for i := 0; i < n; i++ {
            node := queue[i]
            level[i] = node.Val
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        res = append(res, level)
        queue = queue[n:]
    }
    return res
}