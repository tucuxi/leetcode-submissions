/**
 * Definition for singly-linked list.
 * class ListNode(var _x: Int = 0) {
 *   var next: ListNode = null
 *   var x: Int = _x
 * }
 */
object Solution {
    def middleNode(head: ListNode): ListNode = {
         select(head, lengthAcc(0, head) / 2)
    }
    def lengthAcc(acc: Int, head: ListNode): Int = {
        if (head == null) acc else lengthAcc(acc + 1, head.next)
    }
    def select(head: ListNode, index: Int): ListNode = {
        if (index == 0) head else select(head.next, index - 1)
    }
}