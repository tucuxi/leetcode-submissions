/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const MOD = 1_000_000_007

func maxProduct(root *TreeNode) int {
    var allSums []int
    var calculateSums func(*TreeNode) int

    calculateSums = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        sum := root.Val + calculateSums(root.Left) + calculateSums(root.Right)
        allSums = append(allSums, sum)
        return sum
    }

    totalSum := calculateSums(root)
    maxProduct := 0
    
    for _, sum := range allSums {
        product := sum * (totalSum - sum)
        maxProduct = max(maxProduct, product)
    }

    return maxProduct % MOD
}