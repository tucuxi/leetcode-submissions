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
    fun recoverTree(root: TreeNode?): Unit {
        var prev: TreeNode? = null
        var p: TreeNode? = null
        var q: TreeNode? = null
        
        fun walk(node: TreeNode?) {
            node?.run {
                walk(left)
                if (prev != null && prev!!.`val` > `val`) {
                    if (p == null) {
                        p = prev
                        q = this
                    } else if (p!!.`val` > `val`) {
                        q = this
                    }
                }
                prev = this
                walk(right)
            }
        }
        
        walk(root)
        val h = p!!.`val`
        p!!.`val` = q!!.`val`
        q!!.`val` = h
    }
}