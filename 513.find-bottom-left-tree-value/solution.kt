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
    fun findBottomLeftValue(root: TreeNode?): Int {
        var leftMost = -1
        if (root != null) {
            var q = listOf<TreeNode>(root)
            while (q.isNotEmpty()) {
                leftMost = q.first().`val`
                q = q.flatMap { listOf(it.left, it.right) }.filterNotNull()
            }
        }
        return leftMost
    }
}