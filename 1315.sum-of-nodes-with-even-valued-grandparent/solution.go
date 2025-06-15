/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
    return traverse(root, nil, nil, 0)
}

func traverse(current, parent, grandParent *TreeNode, acc int) int {
    if current == nil {
        return acc
    }
    if parent != nil && grandParent != nil && grandParent.Val % 2 == 0 {
        acc += current.Val
    }
    acc = traverse(current.Left, current, parent, acc)
    acc = traverse(current.Right, current, parent, acc)
    return acc
}