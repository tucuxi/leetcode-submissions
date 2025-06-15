/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    res, _ := dfs(root, 0)
    return res
}

func dfs(root *TreeNode, depth int) (*TreeNode, int) {
    if root == nil {
        return root, depth
    } 
    ln, ld := dfs(root.Left, depth+1)
    rn, rd := dfs(root.Right, depth+1)
    if ld > rd {
        return ln, ld+1
    }
    if ld < rd {
        return rn, rd+1
    } 
    return root, ld+1
}