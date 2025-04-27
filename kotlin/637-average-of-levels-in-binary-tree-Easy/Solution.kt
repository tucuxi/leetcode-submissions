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
    fun averageOfLevels(root: TreeNode?): DoubleArray {
        val average = ArrayList<Double>()
        val count = ArrayList<Int>()
        
        fun traverse(level: Int, node: TreeNode?) {
            if (node == null) return
            if (level == average.size) {
                average.add(node.`val`.toDouble())
                count.add(1)
            } else {
                average[level] = ((average[level] * count[level]) + node.`val`) / ++count[level]  
            }
            traverse(level + 1, node.left)
            traverse(level + 1, node.right)
        }
        
        traverse(0, root)
        return average.toDoubleArray()
    }
}