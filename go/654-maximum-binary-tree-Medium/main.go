/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    k := 0
    max := math.MinInt
    for i, n := range nums {
        if n > max {
            max = n
            k = i
        }
    }
    return &TreeNode{
        Val: nums[k],
        Left: constructMaximumBinaryTree(nums[:k]),
        Right: constructMaximumBinaryTree(nums[k + 1:]),
    }
}