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
    fun deepestLeavesSum(root: TreeNode?): Int {
        if (root == null) return 0
        var sum = 0
        val q = ArrayDeque<TreeNode>()
        q.add(root)
        while (q.isNotEmpty()) {
            sum = 0
            repeat(q.size) {
                with(q.poll()) {
                    sum += `val`
                    if (left != null) {
                        q.offer(left)
                    }
                    if (right != null) {
                        q.offer(right)
                    }
                }
            }
        }
        return sum
    }
}