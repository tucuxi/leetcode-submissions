/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
    checkVal := []func(int) bool {
        func(a int) bool { return a % 2 == 1 },
        func(a int) bool { return a % 2 == 0 },
    }
    checkOrder := []func(int, int) bool {
        func(a, b int) bool { return a < b },
        func(a, b int) bool { return a > b },
    }
    q := []*TreeNode{root}
    l := 0
    for (len(q) > 0) {
        var p []*TreeNode
        for i := range q {
            if !checkVal[l](q[i].Val) {
                return false
            }
            if i > 0 && !checkOrder[l](q[i-1].Val, q[i].Val) {
                return false
            }
            if q[i].Left != nil {
                p = append(p, q[i].Left)
            }
            if q[i].Right != nil {
                p = append(p, q[i].Right)
            }
        }
        q = p
        l = 1 - l
    }
    return true
}