/**
 * Definition for a binary tree node.
 * class TreeNode(var _value: Int) {
 *   var value: Int = _value
 *   var left: TreeNode = null
 *   var right: TreeNode = null
 * }
 */
object Solution {
    def hasPathSum(root: TreeNode, sum: Int): Boolean = {
        def hasPathSumAcc(acc: Int, node: TreeNode): Boolean = {
            if (node == null)
                false
            else if (node.left == null && node.right == null)
                acc + node.value == sum
            else 
                hasPathSumAcc(acc + node.value, node.left) || 
                hasPathSumAcc(acc + node.value, node.right)
        }
        hasPathSumAcc(0, root)
    }
}