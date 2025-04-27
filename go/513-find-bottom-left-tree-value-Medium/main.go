/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    leftMost := -1
    q := []*TreeNode{root}
    for len(q) > 0 {
        p := make([]*TreeNode, len(q))
        copy(p, q)
        q = q[:0]
        leftMost = p[0].Val
        for _, v := range p {
            if v.Left != nil {
                q = append(q, v.Left)
            }
            if v.Right != nil {
                q = append(q, v.Right)
            }
        }
    }
    return leftMost
}

