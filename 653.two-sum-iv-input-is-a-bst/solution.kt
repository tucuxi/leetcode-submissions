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
    fun findTarget(root: TreeNode?, k: Int): Boolean {
        val values = mutableSetOf<Int>()
        
        fun walk(node: TreeNode?): Boolean {
            if (node == null) {
                return false
            }
            if (values.contains(k - node.`val`)) {
                return true
            }
            values.add(node.`val`)
            return walk(node.left) || walk(node.right)
        }
        
        return walk(root)
    }
}