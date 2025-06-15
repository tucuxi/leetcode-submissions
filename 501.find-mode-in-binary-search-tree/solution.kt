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
    fun findMode(root: TreeNode?): IntArray {
        val f = HashMap<Int, Int>()
        count(root, f)
        val max = f.values.max() ?: 0
        return f.filter { (_, v) -> v == max }.map { (k, _) -> k }.toIntArray()
    }
    
    fun count(root: TreeNode?, f: MutableMap<Int, Int>) {
        if (root != null) {
            f.put(root.`val`, (f[root.`val`] ?: 0) + 1)
            count(root.left, f)
            count(root.right, f)
        }
    }
}