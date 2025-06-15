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
    fun increasingBST(root: TreeNode?): TreeNode? {
        var p: TreeNode? = null
        
        fun walk(node: TreeNode?) {
            node?.run {
                walk(right)
                right = p
                p = this
                walk(left)
                left = null
            }
        }
        
        walk(root)
        return p       
    }
}