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
    fun levelOrder(root: TreeNode?): List<List<Int>> {
        val res: MutableList<List<Int>> = mutableListOf()
        val queue = ArrayDeque<TreeNode>()
        queue.offerUnlessNull(root)
        while (queue.isNotEmpty()) {
            val level = mutableListOf<Int>()
            repeat(queue.size) {
                val node = queue.poll()
                level.add(node.`val`)
                queue.offerUnlessNull(node.left)
                queue.offerUnlessNull(node.right)
            }
            res.add(level)
        }
        return res
    }
}

inline fun <T> ArrayDeque<T>.offerUnlessNull(elem: T?) {
    if (elem != null) offer(elem)
}