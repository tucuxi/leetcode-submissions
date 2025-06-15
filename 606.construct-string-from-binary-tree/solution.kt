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
    fun tree2str(root: TreeNode?): String {
        val res = StringBuilder()
        
        fun walk(node: TreeNode) {
            with(node) {
                res.append(`val`)
                if (left != null) {
                    res.append('(')
                    walk(left)
                    res.append(')')
                } else if (right != null) {
                    res.append("()")
                }
                if (right != null) {
                    res.append('(')
                    walk(right)
                    res.append(')')
                }               
            }
        }
        
        if (root != null) {
            walk(root)
        }
        return res.toString()
    }
}