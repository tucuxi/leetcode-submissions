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
    fun sumRootToLeaf(root: TreeNode?): Int {
        var total = 0
        
        fun traverse(acc: Int, node: TreeNode?) {
            if (node == null) return            
            val newAcc = 2 * acc + node.`val`
            if (isLeaf(node)) {
                total += newAcc   
            } else {
                traverse(newAcc, node.left)
                traverse(newAcc, node.right)
            }
        }
        
        traverse(0, root)
        return total
    }
    
    private fun isLeaf(node: TreeNode) =
        node.left == null && node.right == null
}