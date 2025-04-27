/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root != nil {
        depth := 0
        for q := []*TreeNode{root}; len(q) > 0; {
            depth++
            k := len(q)
            for i := 0; i < k; i++ {
                node := q[i]
                if node.Left == nil && node.Right == nil {
                    return depth
                }
                if node.Left != nil {
                    q = append(q, node.Left)
                }
                if node.Right != nil {
                    q = append(q, node.Right)
                }
            }
            q = q[k:]
        }
    }
    return 0
}