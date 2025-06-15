/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(traversal string) *TreeNode {
    var t []*TreeNode
    b := []byte(traversal)

    for len(b) > 0 {
        d := 0
        for b[0] == '-' {
            d++
            b = b[1:]
        }

        v := 0
        for len(b) > 0 && b[0] >= '0' && b[0] <= '9' {
            v = 10 * v + int(b[0] - '0')
            b = b[1:]
        }

        n := &TreeNode{Val: v}

        if d == len(t) {
            t = append(t, n)
        } else {
            t[d] = n
        }

        if d > 0 {
            p := t[d - 1]
            if p.Left == nil {
                p.Left = n
            } else {
                p.Right = n
            }
        }
    }

    return t[0]
}

