/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    var res []int
    
    if root != nil {
        for q := []*TreeNode{root}; len(q) > 0; {
            levelMax := math.MinInt
            for range len(q) {
                levelMax = max(levelMax, q[0].Val)
                if q[0].Left != nil {
                    q = append(q, q[0].Left)
                }
                if q[0].Right != nil {
                    q = append(q, q[0].Right)
                }
                q = q[1:]
            }
            res = append(res, levelMax)
        }
    }

    return res
}