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
    fun isUnivalTree(root: TreeNode?): Boolean {
        if (root == null) return true
        val value = root.`val`
        
        fun match(node: TreeNode?): Boolean =
            (node == null) || 
            (node.`val` == value && match(node.left) && match(node.right))
            
        return match(root)
    }
}