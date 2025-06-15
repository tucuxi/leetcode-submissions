/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    k := index(inorder, preorder[0])
    return &TreeNode{
        preorder[0],
        buildTree(preorder[1:k+1], inorder[:k]),
        buildTree(preorder[k+1:], inorder[k+1:])}
}

func index(s []int, val int) int {
    for i := range s {
        if s[i] == val {
            return i
        }
    }
    return len(s)
}