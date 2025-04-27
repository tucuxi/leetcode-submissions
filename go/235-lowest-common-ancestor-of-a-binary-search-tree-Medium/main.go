/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var lo, hi int
    if p.Val < q.Val {
        lo = p.Val
        hi = q.Val
    } else {
        lo = q.Val
        hi = p.Val
    }
    for root != nil {
        if root.Val < lo {
            root = root.Right
        } else if root.Val > hi {
            root = root.Left
        } else {
            return root
        }
    }
    return nil
}