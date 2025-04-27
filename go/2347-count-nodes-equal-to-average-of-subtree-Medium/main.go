/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) int {
    res := 0
    
    var traverse func(*TreeNode) (int, int)
    traverse = func(root *TreeNode) (int, int) {
        if root == nil {
            return 0, 0
        }
        numL, sumL := traverse(root.Left)
        numR, sumR := traverse(root.Right)
        num := numL + numR + 1
        sum := sumL + sumR + root.Val
        if sum / num == root.Val {
            res++
        }
        return num, sum
    }
    
    traverse(root)
    return res
}