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
    fun mergeKLists(lists: Array<ListNode?>): ListNode? {
        val h = ListNode(0)
        var p = h
        while (true) {
            var minValue = Int.MAX_VALUE
            var minIndex = -1
            lists.forEachIndexed { index, list ->
                if (list != null && list.`val` < minValue) {
                    minValue = list.`val`
                    minIndex = index
                }
            }
            if (minIndex == -1) {
                return h.next
            }
            p.next = lists[minIndex]
            lists[minIndex] = lists[minIndex]!!.next
            p = p.next
        }
    }
}