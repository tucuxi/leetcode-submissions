/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    q := []*TreeNode{root}
    for d := 0;; d++ {
        for i := len(q); i > 0; i-- {
            if  p := q[0]; p != nil {
                q = append(q, p.Left)
                q = append(q, p.Right)
            }
            q = q[1:]
        }
        if len(q) == 0 {
            return d
        }
    }
}
