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
    siblingVal int
}

func replaceValueInTree(root *TreeNode) *TreeNode {
    q := []entry{{root, 0}}
    for len(q) > 0 {
        k := len(q)
        s := 0
        for i := 0; i < k; i++ {
            s += q[i].node.Val
        }
        for i := 0; i < k; i++ {
            u := q[i].node
            u.Val = s - u.Val - q[i].siblingVal
            if u.Left != nil {
                rv := 0
                if u.Right != nil {
                    rv = u.Right.Val
                }
                q = append(q, entry{u.Left, rv})
            }
            if u.Right != nil {
                lv := 0
                if u.Left != nil {
                    lv = u.Left.Val
                }
                q = append(q, entry{u.Right, lv})
            }
        }
        q = q[k:]
    }
    return root
}