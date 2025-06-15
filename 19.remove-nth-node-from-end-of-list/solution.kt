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
    fun removeNthFromEnd(head: ListNode?, n: Int): ListNode? {
        var q = head
        repeat(n) { q = q?.next }
        if (q == null) {
            return head!!.next
        }
        var p = head
        while (q!!.next != null) {
            p = p!!.next
            q = q!!.next
        }
        p!!.next = p!!.next!!.next
        return head
    }
}