/**
 * Definition for a binary tree node.
 * class TreeNode(var _value: Int) {
 *   var value: Int = _value
 *   var left: TreeNode = null
 *   var right: TreeNode = null
 * }
 */
object Solution {
    def rangeSumBST(root: TreeNode, L: Int, R: Int): Int = {
        def rangeSum(node: TreeNode): Int = {
            if (node == null) 
                0
            else if (node.value < L)
                rangeSum(node.right)
            else if (node.value > R)
                rangeSum(node.left)
            else
                rangeSum(node.left) + node.value + rangeSum(node.right)
        }
        rangeSum(root)
    }
}
