/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if root == nil {
        return nil
    }
    if depth == 1 {
        return &TreeNode{val, root, nil}
    }
    if depth == 2 { 
        root.Left = &TreeNode{val, root.Left, nil}
        root.Right = &TreeNode{val, nil, root.Right}
    } else {
        addOneRow(root.Left, val, depth - 1)
        addOneRow(root.Right, val, depth - 1)
    }
    return root
}