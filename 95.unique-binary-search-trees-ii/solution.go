/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    return generate(1, n) 
}

func generate(l, r int) []*TreeNode {
    if l > r {
        return []*TreeNode{nil}
    }
    var res []*TreeNode
    for i := l; i <= r; i++ {
        for _, subL := range generate(l, i - 1) {
            for _, subR := range generate(i + 1, r) {
                root := &TreeNode{i, subL, subR}
                res = append(res, root)
            }
        }
    } 
    return res
}