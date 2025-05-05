/**
 * Definition for singly-linked list.
 * class ListNode(_x: Int = 0, _next: ListNode = null) {
 *   var next: ListNode = _next
 *   var x: Int = _x
 * }
 */
object Solution {
    def mergeTwoLists(l1: ListNode, l2: ListNode): ListNode = {
        if (l1 == null) return l2
        if (l2 == null) return l1
        var pred: ListNode = null
        var p1 = l1
        var p2 = l2
        if (p1.x <= p2.x) {
            pred = p1
            p1 = p1.next
        } else {
            pred = p2
            p2 = p2.next
        }
        val head = pred
        while (p1 != null && p2 != null) {
            if (p1.x <= p2.x) {
                pred.next = p1
                pred = p1
                p1 = p1.next
            } else {
                pred.next = p2
                pred = p2
                p2 = p2.next
            }
        }
        pred.next = if (p1 != null) p1 else p2
        head
    }
}