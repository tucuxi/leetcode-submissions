/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimumOperations(root *TreeNode) int {
    res := 0
    q := []*TreeNode{root}
    for len(q) > 0 {
        n := len(q)
        target := make([]int, n)
        pos := make(map[int]int)
        for i := range n {
            target[i] = q[i].Val
            pos[q[i].Val] = i
            if q[i].Left != nil {
                q = append(q, q[i].Left)
            }
            if q[i].Right != nil {
                q = append(q, q[i].Right)
            }
        }
        slices.Sort(target)
        for i := range n {
            if q[i].Val != target[i] {
                res++
                currentPos := pos[target[i]]
                pos[q[i].Val] = currentPos
                q[currentPos].Val = q[i].Val
            }
        }
        q = q[n:]
    }
    return res
}