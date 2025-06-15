/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    var minDiff = math.MaxInt
    var prev *TreeNode
    var inorder func(*TreeNode)
    
    inorder = func(root *TreeNode) {
        if root != nil {
            inorder(root.Left)
            if prev != nil {
                if d := root.Val - prev.Val; d < minDiff {
                    minDiff = d
                }
            }
            prev = root
            inorder(root.Right)
        }
    }
    
    inorder(root)    
    return minDiff
}