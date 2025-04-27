/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
    var res int
    var f func(*TreeNode) (int, bool)
    
    f = func(root *TreeNode) (int, bool) {
        if root == nil {
            return 0, false
        }
        t1, found1 := f(root.Left)
        t2, found2 := f(root.Right)
        if root.Val == start {
            t := max(t1, t2)
            res = max(res, t)
            return 1, true
        }
        if found1 {
            res = max(res, t1+t2)
            return t1+1, true
        }
        if found2 {
            res = max(res, t1+t2)
            return t2+1, true
        }
        return max(t1, t2) + 1, false
    }
    
    f(root)
    return res
}
    
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}