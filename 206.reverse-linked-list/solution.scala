/**
 * Definition for singly-linked list.
 * class ListNode(var _x: Int = 0) {
 *   var next: ListNode = null
 *   var x: Int = _x
 * }
 */
object Solution {
    def reverseList(head: ListNode): ListNode = {
        def reverseListRec(prev: ListNode, curr: ListNode): ListNode = {
            if (curr != null) {
                val next = curr.next
                curr.next = prev
                reverseListRec(curr, next)
             } else {
                prev
             }
        }
        reverseListRec(null, head)
    }
}