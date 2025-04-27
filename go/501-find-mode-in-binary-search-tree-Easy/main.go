/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
    curVal := math.MinInt
    curLen := 0
    maxVal := []int{}
    maxLen := curLen
    
    var rec func(*TreeNode)
    rec = func(node *TreeNode) {
        if node != nil {
            rec(node.Left)
            if node.Val == curVal {
                curLen++
            } else {
                curVal = node.Val
                curLen = 1
            }
            if curLen == maxLen {
                maxVal = append(maxVal, node.Val)
            } else if curLen > maxLen {
                maxLen = curLen
                maxVal = []int{node.Val}
            }
            rec(node.Right)
        }
    }
    
    rec(root)
    return maxVal
}