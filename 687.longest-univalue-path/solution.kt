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
    fun longestUnivaluePath(root: TreeNode?): Int {
        var res = 0

        fun dfs(node: TreeNode?, parentValue: Int): Int {
            if (node == null) return 0
            val left = dfs(node.left, node.`val`)
            val right = dfs(node.right, node.`val`)
            res = maxOf(res, left + right)
            return if (node.`val` == parentValue) maxOf(left, right) + 1 else 0
        }

        dfs(root, -1)
        return res
    }

}