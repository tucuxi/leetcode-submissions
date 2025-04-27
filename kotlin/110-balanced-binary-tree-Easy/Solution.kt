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
    fun isBalanced(root: TreeNode?): Boolean {
        var bal = true

        fun height(node: TreeNode?): Int = 
            if (node == null)
                -1
             else {
                 val lh = height(node.left)
                 val rh = height(node.right)
                 if (Math.abs(lh - rh) > 1)
                    bal = false
                 maxOf(lh, rh) + 1
             }

        height(root)
        return bal
    }
}