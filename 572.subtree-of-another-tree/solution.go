/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if subRoot == nil {
        return true
    }
    if root == nil {
        return false
    }
    return equal(root, subRoot) ||
        isSubtree(root.Left, subRoot) ||
        isSubtree(root.Right, subRoot)
}

func equal(node1, node2 *TreeNode) bool {
    if node1 == nil {
        return node2 == nil
    }
    if node2 == nil {
        return node1 == nil
    }
    return node1.Val == node2.Val &&
        equal(node1.Left, node2.Left) &&
        equal(node1.Right, node2.Right)
}