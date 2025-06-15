/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    res := [][]int{}
    q := []*TreeNode{}
    if root != nil {
        q = append(q, root)
    }
    for len(q) > 0 {
        level := make([]int, 0, len(q))
        for i := len(q); i > 0; i-- {
            n := q[0]
            q = q[1:]
            level = append(level, n.Val)
            if n.Left != nil {
                q = append(q, n.Left)
            }
            if n.Right != nil {
                q = append(q, n.Right)
            }
        }
        res = append([][]int{level}, res...)
    }
    return res    
}
