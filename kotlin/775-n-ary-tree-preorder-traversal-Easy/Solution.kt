/**
 * Definition for a Node.
 * class Node(var `val`: Int) {
 *     var children: List<Node?> = listOf()
 * }
 */

class Solution {
    fun preorder(root: Node?): List<Int> {
        val res = mutableListOf<Int>()
        
        fun rec(node: Node?) {
            node ?: return
            res.add(node.`val`)
            node.children.forEach { rec(it) }
        }
                
        rec(root)
        return res
    }
}