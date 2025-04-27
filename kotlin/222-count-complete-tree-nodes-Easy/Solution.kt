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
    fun countNodes(root: TreeNode?): Int {
        val dl = depth(root) { it.left }
        val dr = depth(root) { it.right }
        return if (dl == dr) {
            (1 shl dl) - 1
        } else {
            countNodes(root!!.left) + countNodes(root!!.right) + 1
        }
    }
    
    fun depth(root: TreeNode?, step: (node: TreeNode) -> TreeNode?): Int {
        var node = root
        var depth = 0
        while (node != null) {
            depth++
            node = step(node)
        }
        return depth
    }
}