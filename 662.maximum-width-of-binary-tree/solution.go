/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type entry struct {
    node *TreeNode
    index int
}

func widthOfBinaryTree(root *TreeNode) int {
    res := 0
    q := []entry{{root, 0}} 
    for len(q) > 0 {
        k := len(q)
        if w := q[k - 1].index - q[0].index + 1; w > res {
            res = w
        }
        for i := 0; i < k; i++ {
            if next := q[i].node.Left; next != nil {
                q = append(q, entry{next, 2 * q[i].index + 1})
            }
            if next := q[i].node.Right; next != nil {
                q = append(q, entry{next, 2 * q[i].index + 2})
            }
        }
        q = q[k:]
    }
    return res
}
