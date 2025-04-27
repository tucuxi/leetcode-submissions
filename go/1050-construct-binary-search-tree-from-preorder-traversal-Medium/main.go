/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    k := 1 + sort.Search(len(preorder) - 1, func(i int) bool {
        return preorder[i + 1] > preorder[0]
    })
    return &TreeNode{
        Val: preorder[0],
        Left: bstFromPreorder(preorder[1:k]),
        Right: bstFromPreorder(preorder[k:]),
    }
}