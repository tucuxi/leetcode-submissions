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
    fun minDiffInBST(root: TreeNode?): Int {
        val inorder = root?.toList() ?: emptyList()
        var min = Int.MAX_VALUE
        for (i in 1..inorder.lastIndex) {
            min = minOf(min, inorder[i] - inorder[i - 1])            
        }
        return min
    }
}

private fun TreeNode.toList(): List<Int> {
    val res = mutableListOf<Int>()
    rec(this, res)
    return res
}    

private fun rec(root: TreeNode?, res: MutableList<Int>) {
    root?.run {
        rec(left, res)
        res.add(`val`)
        rec(right, res)
    }
}
