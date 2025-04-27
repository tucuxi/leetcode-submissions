/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countPairs(root *TreeNode, distance int) int {
    goodPairs := 0

    var dfs func(*TreeNode) []int
    dfs = func(node *TreeNode) []int {
        if node == nil {
            return nil
        }
        if node.Left == nil && node.Right == nil {
            return []int{0}
        }

        left := dfs(node.Left) 
        right := dfs(node.Right)
        
        for _, l := range left {
            for _, r := range right {
                if l + r + 2 <= distance {
                    goodPairs++
                }
            }
        }

        var leaveDistances []int
        
        for _, l := range left {
            leaveDistances = append(leaveDistances, l + 1)
        }
        for _, r := range right {
            leaveDistances = append(leaveDistances, r + 1)
        }

        return leaveDistances
    }

    dfs(root)
    return goodPairs
}