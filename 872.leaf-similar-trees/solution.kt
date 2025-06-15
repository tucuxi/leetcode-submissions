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
    fun leafSimilar(root1: TreeNode?, root2: TreeNode?): Boolean {
        return leaves(root1) == leaves(root2)
    }
    
    fun leaves(root: TreeNode?): List<Int> {
        val res = mutableListOf<Int>()
        
        fun walkTree(root: TreeNode) {
            if (root.left == null && root.right == null) {
                res.add(root.`val`)
                return
            }
            root.left?.run { walkTree(this) }
            root.right?.run { walkTree(this) }
        }
    
        root?.run { walkTree(this) }
        return res
    }
}