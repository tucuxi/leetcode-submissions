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
class Solution {
    fun maxLevelSum(root: TreeNode?): Int {
        var maxSum = Int.MIN_VALUE
        var maxLevel = 0
        var currentLevel = 1
        var nodes = listOfNotNull(root)

        while (nodes.isNotEmpty()) {
            val currentSum = nodes.sumOf { it.`val` }
            if (currentSum > maxSum) {
                maxSum = currentSum
                maxLevel = currentLevel
            }
            currentLevel++
            nodes = nodes.flatMap { listOfNotNull(it.left, it.right) }
        }
        return maxLevel
    }
}