/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    return build(inorder(root))
}

func inorder(root *TreeNode) []int {
    var sorted []int
    var traverse func(*TreeNode)
    
    traverse = func(root *TreeNode) {
        if root != nil {
            traverse(root.Left)
            sorted = append(sorted, root.Val)
            traverse(root.Right)
        }
    }
    traverse(root)
    return sorted
}

func build(sorted []int) *TreeNode {
    if len(sorted) == 0 {
        return nil
    }
    m := len(sorted) / 2
    return &TreeNode{sorted[m], build(sorted[:m]), build(sorted[m + 1:])}
}