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
    fun isPalindrome(head: ListNode?): Boolean {
        val forward = StringBuilder()
        val backward = StringBuilder()
        var node = head
        while (node != null) {
            forward.append(node.`val`)
            backward.insert(0, node.`val`)
            node = node.next
        }
        return forward.toString() == backward.toString()
    }
}