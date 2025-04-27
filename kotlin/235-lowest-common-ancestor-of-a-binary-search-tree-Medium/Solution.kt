/**
 * Definition for a binary tree node.
 * class TreeNode(var `val`: Int = 0) {
 *     var left: TreeNode? = null
 *     var right: TreeNode? = null
 * }
 */

class Solution {
    fun lowestCommonAncestor(root: TreeNode?, p: TreeNode?, q: TreeNode?): TreeNode? {
        val lo = minOf(p!!.`val`, q!!.`val`)
        val hi = maxOf(p!!.`val`, q!!.`val`)
        var node = root
        
        while (node != null) {
            node = when {
                node.`val` > hi -> node.left
                node.`val` < lo -> node.right
                else -> return node
            }
        } 
        return node
    }
}