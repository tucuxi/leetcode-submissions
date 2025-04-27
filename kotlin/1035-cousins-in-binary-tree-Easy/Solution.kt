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
    fun isCousins(root: TreeNode?, x: Int, y: Int): Boolean {
        var px = 0
        var dx = 0
        var py = 0
        var dy = 0
        
        fun dfs(node: TreeNode?, p: Int, d: Int) {
            node?.run {
                when(`val`) {
                    x -> {
                        px = p
                        dx = d
                    }
                    y -> {
                        py = p
                        dy = d
                    }
                }
                dfs(left, `val`, d + 1)
                dfs(right, `val`, d + 1)
            }
        }
        
        dfs(root, 0, 0)
        return dx == dy && px != py
    }   
}