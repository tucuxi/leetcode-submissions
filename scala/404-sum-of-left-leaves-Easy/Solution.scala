/**
 * Definition for a binary tree node.
 * class TreeNode(var _value: Int) {
 *   var value: Int = _value
 *   var left: TreeNode = null
 *   var right: TreeNode = null
 * }
 */
object Solution {
    def sumOfLeftLeaves(root: TreeNode): Int = {
        def sum(node: TreeNode, isLeft: Boolean): Int = {
            if (node == null) 0
            else if (isLeft && isLeaf(node)) node.value
            else sum(node.left, true) + sum(node.right, false)
        }
        def isLeaf(node: TreeNode) =
            node != null && node.left == null && node.right == null
        sum(root, false)
    }
}