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
    fun sortedArrayToBST(nums: IntArray): TreeNode? {
        fun rec(lo: Int, hi: Int): TreeNode? {
            if (lo > hi) return null
            val m = lo + (hi - lo + 1) / 2
            val ti = TreeNode(nums[m])
            ti.left = rec(lo, m - 1)
            ti.right = rec(m + 1, hi)
            return ti
        }
        
        return rec(0, nums.size - 1)
    }
}