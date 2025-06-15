import kotlin.math.min

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
    fun minDepth(root: TreeNode?) =
        if (root == null) {
            0
        } else {
            descend(root, 1)
        }
    
    fun descend(node: TreeNode, depth: Int): Int =
        if (node.left == null && node.right == null) {
            depth
        } else if (node.left == null) {
            descend(node.right, depth + 1)
        } else if (node.right == null) {
            descend(node.left, depth + 1) 
        } else {
            val leftDepth = descend(node.left, depth + 1)
            val rightDepth = descend(node.right, depth + 1)
            min(leftDepth, rightDepth)
        }
}