/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    res := 0
    
    var traverse func(*TreeNode, int)
    traverse = func(root *TreeNode, sum int) {
        if root == nil {
            return
        }
        sum = 10 * sum + root.Val
        if root.Left == nil && root.Right == nil {
            res += sum
            return
        }
        traverse(root.Left, sum)
        traverse(root.Right, sum)
    }
    
    traverse(root, 0)
    return res
}