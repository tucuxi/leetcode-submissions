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
    fun findSecondMinimumValue(root: TreeNode?): Int {
        return root?.run {
            if (left == null || right == null) {
                -1
            } else {
                val lv = if (`val` == left.`val`) {
                    findSecondMinimumValue(root.left)
                } else {
                    left.`val`
                }
                val rv = if (`val` == right.`val`) {
                    findSecondMinimumValue(root.right)
                } else {
                    right.`val`
                }
                if (lv != -1 && rv != -1) {
                    minOf(lv, rv)
                } else if (lv != -1) {
                    lv
                } else {
                    rv
                }
            }
        } ?: -1
    }
}