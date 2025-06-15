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
    fun kthSmallest(root: TreeNode?, k: Int): Int {
        var i = 0
        var res = 0
        
        fun walk(node: TreeNode?): Unit {
            if (i < k) {
                node?.run {
                    walk(left)
                    if (++i == k) {
                        res = `val`
                    }
                    walk(right)
                }
            }
        }
        
        walk(root)
        return res
    }
}