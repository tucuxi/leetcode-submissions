/**
 * Definition for singly-linked list.
 * class ListNode(var _x: Int = 0) {
 *   var next: ListNode = null
 *   var x: Int = _x
 * }
 */
object Solution {
    def deleteDuplicates(head: ListNode): ListNode = {
        def del(head: ListNode, prev: ListNode, curr: ListNode): ListNode = {
            if (curr == null) return head
            if (curr.x == prev.x) {
                prev.next = curr.next
                del(head, prev, curr.next)
            } else
                del(head, curr, curr.next)
        }
        if (head == null) return null
        del(head, head, head.next)
    }
}