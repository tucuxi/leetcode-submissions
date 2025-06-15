/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    sum := robRec(root)
    return max(sum...)
}

func robRec(root *TreeNode) []int {
    if root == nil {
        return []int{0, 0}
    }
    left := robRec(root.Left)
    right := robRec(root.Right)
    return []int{max(left[0], left[1]) + max(right[0], right[1]), root.Val + left[0] + right[0]}
}

func max(a ...int) int {
    m := 0
    for _, v := range a {
        if v > m {
            m = v
        }
    }
    return m
}