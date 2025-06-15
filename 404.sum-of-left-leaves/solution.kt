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
    fun sumOfLeftLeaves(root: TreeNode?) = sum(root, false)
    
    private fun sum(node: TreeNode?, isLeft: Boolean): Int = 
        if (node == null) {
            0
        } else {
            with (node) {
                if (left == null && right == null && isLeft) {
                    `val`
                } else {
                    sum(left, true) + sum(right, false)
                }
            }
        }
}