/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    var (
        deleteSet [1001]bool
        res []*TreeNode
        f func(*TreeNode, bool) bool
    )

    for _, value := range to_delete {
        deleteSet[value] = true
    }

    f = func(node *TreeNode, isParentDeleted bool) bool {
        if node == nil {
            return false
        }
        isDeleted := deleteSet[node.Val]
        if f(node.Left, isDeleted) {
            node.Left = nil
        }
        if f(node.Right, isDeleted) {
            node.Right = nil
        }
        if !isDeleted && isParentDeleted {
            res = append(res, node)
        }
        return isDeleted
    }

    f(root, true)
    return res
}