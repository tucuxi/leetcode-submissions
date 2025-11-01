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
    fun modifiedList(nums: IntArray, head: ListNode?): ListNode? {
        val set = nums.toSet()
        val sentinel = ListNode(0).apply { next = head }
        var p = sentinel

        while (p.next != null) {
            if (p.next.`val` in set) {
                p.next = p.next.next
            } else {
                p = p.next
            }
        }

        return sentinel.next
    }
}