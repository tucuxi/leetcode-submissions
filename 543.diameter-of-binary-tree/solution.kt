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
    fun diameterOfBinaryTree(root: TreeNode?): Int {
        var diameter = 0
        
        fun recurse(node: TreeNode?): Int {
            if (node == null) {
                return 0
            }
            val left = recurse(node.left)
            val right = recurse(node.right)
            diameter = Math.max(diameter, left + right)
            return Math.max(left, right) + 1
        }
        
        recurse(root)
        return diameter
    }
}