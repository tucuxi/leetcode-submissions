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
    fun pathSum(root: TreeNode?, targetSum: Int): List<List<Int>> {
        val res = mutableListOf<List<Int>>()
        
        fun TreeNode.walk(rest: Int, path: List<Int>) {
            val newRest = rest - `val`
            val newPath = path + `val`
            if (left == null && right == null) {
                if (newRest == 0) {
                    res.add(newPath)
                }
                return
            }
            left?.walk(newRest, newPath)
            right?.walk(newRest, newPath)
        }
        
        root?.walk(targetSum, emptyList<Int>())
        return res
    }
}