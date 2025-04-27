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
    fun addTwoNumbers(l1: ListNode?, l2: ListNode?): ListNode? {
        val headSum = (l1?.`val` ?: 0) + (l2?.`val` ?: 0)
        val headNode = ListNode(headSum % 10)
        var prevNode = headNode
        var carry = headSum / 10
        var itr1 = l1
        var itr2 = l2
        while (true) {
            itr1 = itr1?.next ?: null
            itr2 = itr2?.next ?: null
            if (itr1 == null && itr2 == null && carry == 0) break
            val nextSum = (itr1?.`val` ?: 0) + (itr2?.`val` ?: 0) + carry
            val nextNode = ListNode(nextSum % 10)
            prevNode.next = nextNode
            prevNode = nextNode
            carry = nextSum / 10
        }
        return headNode
    }
}