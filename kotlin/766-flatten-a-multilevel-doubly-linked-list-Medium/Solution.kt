/**
 * Definition for a Node.
 * class Node(var `val`: Int) {
 *     var prev: Node? = null
 *     var next: Node? = null
 *     var child: Node? = null
 * }
 */

class Solution {
    fun flatten(root: Node?): Node? {
        var node = root
        while (node != null) {
            var child = node.child
            if (child != null) {
                var childEnd: Node = child
                while (childEnd.next != null) {
                    childEnd = childEnd.next!!
                }
                childEnd.next = node.next
                childEnd.next?.prev = childEnd
                node.next = child
                node.next!!.prev = node
                node.child = null
            }
            node = node.next
        }
        return root
    }
}