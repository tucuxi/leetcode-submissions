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
    fun binaryTreePaths(root: TreeNode?): List<String> {
        val res = mutableListOf<String>()
        
        fun descend(path: Sequence<Int>, node: TreeNode?) {
            if (node == null) return
            if (node.left == null && node.right == null) {
                res += (path + node.`val`).joinToString("->")
                return
            }
            descend(path + node.`val`, node.left)
            descend(path + node.`val`, node.right)
        }
        
        descend(sequenceOf<Int>(), root)
        return res
    }
}