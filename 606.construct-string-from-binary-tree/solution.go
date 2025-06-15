/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
    switch {
    case root == nil:
        return ""    
    case root.Right != nil:
        return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
    case root.Left != nil:
        return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
    default:
        return fmt.Sprintf("%d", root.Val)
    }
}