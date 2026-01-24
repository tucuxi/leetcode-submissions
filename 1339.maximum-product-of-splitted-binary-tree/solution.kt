/**
 * Example:
 * var ti = TreeNode(5)
 * var v = ti.`val`
 * Definition for a binary tree node.
 * class TreeNode(var `val`: Int) {
 *     var left: TreeNode? = null
 *     var right: TreeNode? = null
 * }
 */
const val MOD = 1_000_000_007

class Solution {
    fun maxProduct(root: TreeNode?): Int {
        val allSums = mutableListOf<Long>()

        fun calculateSum(node: TreeNode?): Long {
            if (node == null) return 0L
            val currentSum = node.`val` + calculateSum(node.left) + calculateSum(node.right)
            allSums.add(currentSum)
            return currentSum
        }

        val totalSum = calculateSum(root)
        var maxProduct = 0L

        for (sum in allSums) {
            val product = sum * (totalSum - sum)
            if (product > maxProduct) maxProduct = product
        }

        return (maxProduct % MOD).toInt()
    }
}