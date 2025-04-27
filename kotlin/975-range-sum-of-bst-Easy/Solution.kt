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
    fun rangeSumBST(root: TreeNode?, low: Int, high: Int): Int {
        with(root) {
            return when {
                this == null -> 0 
                `val` < low -> rangeSumBST(right, low, high)
                `val` > high -> rangeSumBST(left, low, high)
                else -> rangeSumBST(left, low, high) + `val` + rangeSumBST(right, low, high)
            }
        }
    }
}