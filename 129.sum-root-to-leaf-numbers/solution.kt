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
    fun sumNumbers(root: TreeNode?): Int {
        var sum = 0
        
        fun walk(node: TreeNode?, num: Int) {
            if (node == null) return
            val newNum = num * 10 + node.`val`
            if (node.left == null && node.right == null) {
                sum += newNum
            } else {
                walk(node.left, newNum)
                walk(node.right, newNum)
            }
        }
        
        walk(root, 0)
        return sum
    }
}