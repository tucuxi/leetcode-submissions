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
    fun preorderTraversal(root: TreeNode?): List<Int> {
        val res = mutableListOf<Int>()
        
        fun recurse(node: TreeNode?): Unit {
            if (node != null) {
                res.add(node.`val`)
                recurse(node.left)
                recurse(node.right)
            }
        }
        
        recurse(root)
        return res
    }
}