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

import kotlin.math.abs

class Solution {
    fun findTilt(root: TreeNode?): Int =
        rec(root).sumOfTilt
    
    data class Result(
        val sumOfVal: Int,
        val sumOfTilt: Int
    )
    
    private fun rec(root: TreeNode?): Result {
        return if (root == null) {
            Result(0, 0)
        } else {
            val leftResult = rec(root.left)
            val rightResult = rec(root.right)
            val tilt = abs(leftResult.sumOfVal - rightResult.sumOfVal)
            Result(leftResult.sumOfVal + rightResult.sumOfVal + root.`val`,
                   leftResult.sumOfTilt + rightResult.sumOfTilt + tilt)
        }
    }
}