/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    first := true
    prev := 0
    diff := math.MaxInt
    inorder(root, func(val int) {
        if !first && val - prev < diff {
            diff = val - prev
        }
        first = false
        prev = val
    })
    return diff
}

func inorder(root *TreeNode, update func(val int)) {
    if root == nil {
        return
    }
    inorder(root.Left, update)
    update(root.Val)
    inorder(root.Right, update)
}