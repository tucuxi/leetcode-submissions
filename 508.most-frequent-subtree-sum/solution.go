/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
    h := make(map[int]int)
    
    var traverse func(*TreeNode) int
    traverse = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        s := root.Val + traverse(root.Left) + traverse(root.Right)
        h[s]++
        return s
    }
    traverse(root)
    
    sums := make([][2]int, 0, len(h))
    for sum, frq := range h {
        sums = append(sums, [2]int{sum, frq})
    }
    sort.Slice(sums, func(i, j int) bool {
        return sums[i][1] > sums[j][1]
    })
    
    res := []int{sums[0][0]}
    for i := 1; i < len(sums) && sums[i][1] == sums[i - 1][1]; i++ {
        res = append(res, sums[i][0])
    }
    
    return res
}