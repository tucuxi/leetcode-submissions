/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
    q := []*TreeNode{root}
    c := false
    for len(q) > 0 {
        k := len(q)
        for i := 0; i < k; i++ {
            if q[i] == nil {
                c = true
            } else if c {
                return false
            } else {
                q = append(q, q[i].Left, q[i].Right)
            }
        }
        q = q[k:]
    }
    return true
}