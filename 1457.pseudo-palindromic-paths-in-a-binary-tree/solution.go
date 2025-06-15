/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pseudoPalindromicPaths (root *TreeNode) int {
    odd := [10]bool{}
    res := 0
    
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        odd[node.Val] = !odd[node.Val]
        if node.Left == nil && node.Right == nil {
            cnt := 0
            for _, o := range odd {
                if o {
                    cnt++
                }
            }
            if cnt <= 1 {
                res++
            }
        } else {
            dfs(node.Left)
            dfs(node.Right)
        }
        odd[node.Val] = !odd[node.Val]
    }
    
    dfs(root)
    return res
}