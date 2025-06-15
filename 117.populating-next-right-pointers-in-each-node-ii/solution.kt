/**
 * Definition for a Node.
 * class Node(var `val`: Int) {
 *     var left: Node? = null
 *     var right: Node? = null
 *     var next: Node? = null
 * }
 */

class Solution {    
    fun connect(root: Node?): Node? =
        root?.apply {
            left?.next = searchLeft(this)
            right?.next = search(next)
            connect(right)
            connect(left)
        }
	
    private fun search(node: Node?): Node? =
        when {
            node == null -> null
            node.left != null -> node.left
            node.right != null -> node.right
            else -> search(node.next)
        }
	
    private fun searchLeft(node: Node?): Node? =
        when {
            node == null -> null
            node.right != null -> node.right
            else -> search(node.next)        
        }
}