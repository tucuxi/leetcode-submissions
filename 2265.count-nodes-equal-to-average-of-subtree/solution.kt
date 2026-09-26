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
    data class Result(val number: Int, val sum: Int, val count: Int)

    fun averageOfSubtree(root: TreeNode?): Int {
        fun dfs(node: TreeNode?): Result {
            if (node == null) {
                return Result(0, 0, 0)
            }
            val l = dfs(node.left)
            val r = dfs(node.right)
            val number = l.number + r.number + 1
            val sum = l.sum + r.sum + node.`val`
            val count = l.count + r.count + if (node.`val` == sum / number) 1 else 0
            return Result(number, sum, count)
        }

        return dfs(root).count
    }
}