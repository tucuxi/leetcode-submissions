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
    fun goodNodes(root: TreeNode?): Int {
        var count = 0
        
        fun walk(node: TreeNode?, maxVal: Int) {
            node?.apply {
                if (`val` >= maxVal) count++
                val newMaxVal = maxOf(maxVal, `val`)
                walk(left, newMaxVal)
                walk(right, newMaxVal)
            }
        }
        
        walk(root, Int.MIN_VALUE)
        return count
        
    }
}