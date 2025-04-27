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
    fun isEvenOddTree(root: TreeNode?): Boolean {
        var level = listOf(root!!)
        var even = true
        while (level.isNotEmpty()) {
            if (even) {
                if (level.any { node -> node.`val` % 2 == 0 } ||
                    level.zipWithNext().any { (a, b) -> a.`val` >= b.`val` }
                ) {
                    return false
                }
            } else {
                if (level.any { node -> node.`val` % 2 == 1 } ||
                    level.zipWithNext().any { (a, b) -> a.`val` <= b.`val` }
                ) {
                    return false
                }
            }
            level = level.flatMap { listOf(it.left, it.right) }.filterNotNull()
            even = !even
        }
        return true
    }
}