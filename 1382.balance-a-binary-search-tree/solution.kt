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
    fun balanceBST(root: TreeNode?): TreeNode? {
        val nums = mutableListOf<Int>()

        fun inorder(root: TreeNode?) {
            root?.let { node ->
                inorder(node.left)
                nums.add(node.`val`)
                inorder(node.right)
            }
        }

        fun build(list: List<Int>): TreeNode? {
            if (list.isEmpty()) {
                return null
            }
            val m = list.size / 2
            return TreeNode(list[m]).apply {
                left = build(list.subList(0, m))
                right = build(list.subList(m + 1, list.size))
            }
        }

        inorder(root)
        return build(nums)
    }
}