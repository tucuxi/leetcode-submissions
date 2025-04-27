/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    avg := []float64{}
    cnt := []float64{}
    
    var walk func(level int, node *TreeNode)
    walk = func(level int, node *TreeNode) {
        if node == nil {
            return
        }
        if level == len(avg) {
            avg = append(avg, float64(node.Val))
            cnt = append(cnt, 1)
        } else {
            avg[level] = (avg[level] * cnt[level] + float64(node.Val)) / (cnt[level] + 1)
            cnt[level]++
        }
        walk(level + 1, node.Left)
        walk(level + 1, node.Right)
    }
    
    walk(0, root)
    return avg
}