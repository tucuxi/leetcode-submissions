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
    fun maxProduct(root: TreeNode?): Int {
        val sumRoot = sumTree(root)
        val totalSum = sumRoot?.`val` ?: 0 
        var maxProd: Long = 0
        
        fun recurse(root: TreeNode) {
            if (root.left != null) {
                maxProd = maxOf(maxProd, root.left.`val` * (totalSum - root.left.`val`).toLong())
                recurse(root.left)
            }
            if (root.right != null) {
                maxProd = maxOf(maxProd, root.right.`val` * (totalSum - root.right.`val`).toLong())
                recurse(root.right)
            }
        }

        if (sumRoot != null) { 
            recurse(sumRoot)
        }
        return (maxProd % 1_000_000_007).toInt()
    }
    
    fun sumTree(root: TreeNode?): TreeNode? {
        if (root == null) return null
        val leftChild = sumTree(root.left)
        val rightChild = sumTree(root.right)
        val sum = (leftChild?.`val` ?: 0) + (rightChild?.`val` ?: 0) + root.`val`
        return TreeNode(sum).apply {
            left = leftChild
            right = rightChild
        }
    } 
}