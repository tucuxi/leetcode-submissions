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
    fun inorderTraversal(root: TreeNode?): List<Int> {
        val res = mutableListOf<Int>()
        fun recurse(node: TreeNode?) {
            node?.apply {
                recurse(left)
                res.add(`val`)
                recurse(right)
            }
        }
        recurse(root)
        return res
    }
}