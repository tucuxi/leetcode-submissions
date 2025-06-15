/**
 * Example:
 * var li = ListNode(5)
 * var v = li.`val`
 * Definition for singly-linked list.
 * class ListNode(var `val`: Int) {
 *     var next: ListNode? = null
 * }
 */

class Solution {
    fun getIntersectionNode(headA:ListNode?, headB:ListNode?):ListNode? {
        val a = collectNodes(headA)
        var node = headB
        while (node != null) {
            if (a.contains(node)) break
            node = node.next
        }
        return node
        
    }
    
    private fun collectNodes(head: ListNode?): Set<ListNode> {
        val res = mutableSetOf<ListNode>()
        var node = head
        while (node != null) {
            res.add(node)
            node = node.next
        }
        return res
    }
}