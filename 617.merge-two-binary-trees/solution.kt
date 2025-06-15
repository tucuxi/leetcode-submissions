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
    fun mergeTrees(root1: TreeNode?, root2: TreeNode?): TreeNode? {
        root1 ?: return root2
        root2 ?: return root1
        return root1.apply {
            `val` += root2.`val`
            left = mergeTrees(root1.left, root2.left)
            right = mergeTrees(root1.right, root2.right)
        }
    }
}