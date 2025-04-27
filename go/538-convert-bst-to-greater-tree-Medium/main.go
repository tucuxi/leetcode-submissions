/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    sum := 0
    
    var rec func(*TreeNode)
    rec = func(root *TreeNode) {
        if root != nil {
            rec(root.Right)
            root.Val += sum
            sum = root.Val
            rec(root.Left)            
        }
    } 
    
    rec(root)
    return root
}