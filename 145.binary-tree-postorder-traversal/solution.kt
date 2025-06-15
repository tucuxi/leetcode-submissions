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
    fun postorderTraversal(root: TreeNode?): List<Int> {
        val res = mutableListOf<Int>()
        
        fun recurse(node: TreeNode?): Unit {
            node?.apply {
                recurse(left)
                recurse(right)
                res.add(`val`)
            }
        }
        
        recurse(root)
        return res
    }
}