/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    res := [][]int{}
    
    var dfs func(node *TreeNode, rest int, path []int)
    dfs = func(node *TreeNode, rest int, path []int) {
        if node == nil {
            return
        }
        rest -= node.Val
        path = append(path, node.Val)
        if node.Left == nil && node.Right == nil {
            if rest == 0 {
                var newPath []int
                newPath = append(newPath, path...)
                res = append(res, newPath)
                return
            }
        }
        dfs(node.Left, rest, path)
        dfs(node.Right, rest, path)
    }
    
    dfs(root, targetSum, []int{})
    return res
}