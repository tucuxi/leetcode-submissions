/**
 * Definition for a binary tree node.
 * class TreeNode(var _value: Int) {
 *   var value: Int = _value
 *   var left: TreeNode = null
 *   var right: TreeNode = null
 * }
 */
object Solution {
    def isSymmetric(root: TreeNode): Boolean = {
        root == null || eq(root.left, root.right) 
    }
    
    def eq(p: TreeNode, q: TreeNode): Boolean = {
        if (p == null && q == null) return true
        if (p == null || q == null) return false
        (p.value == q.value) && eq(p.left, q.right) && eq(p.right, q.left)
    }
}