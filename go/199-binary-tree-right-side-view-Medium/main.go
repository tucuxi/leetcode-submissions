/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    res := []int{}
    
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, level int) {
        if node == nil {
            return
        }
        if level == len(res) {
            res = append(res, node.Val)
        }
        dfs(node.Right, level + 1)
        dfs(node.Left, level + 1)
    }

    dfs(root, 0)
    return res
}