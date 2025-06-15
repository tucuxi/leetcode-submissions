/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    res, _ := dfs(root, 0)
    return res
}

func dfs(root *TreeNode, depth int) (*TreeNode, int) {
    if root == nil {
        return root, depth
    }
    rootL, depthL := dfs(root.Left, depth+1)
    rootR, depthR := dfs(root.Right, depth+1)
    switch {
    case depthL < depthR: return rootR, depthR+1
    case depthL > depthR: return rootL, depthL+1
    default:              return root, depthL+1
    }
}