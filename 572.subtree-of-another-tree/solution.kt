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
    fun isSubtree(root: TreeNode?, subRoot: TreeNode?): Boolean =
        if (root == null) {
            subRoot == null
        } else {
            equal(root, subRoot) || isSubtree(root.left, subRoot) || isSubtree(root.right, subRoot)
        }
    
    private fun equal(root1: TreeNode?, root2: TreeNode?): Boolean =
        if (root1 == null) {
            root2 == null
        } else {
            root2 != null && root1.`val` == root2.`val` && equal(root1.left, root2.left) && equal(root1.right, root2.right)
        }
}
