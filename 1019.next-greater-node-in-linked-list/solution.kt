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
    fun nextLargerNodes(head: ListNode?): IntArray {
        val res = mutableListOf<Pair<Int, Int>>()
        val stack = mutableListOf<Pair<Int, Int>>()
        
        var i = 0
        var p = head

        while (p != null) {
            while (stack.isNotEmpty() && stack.last().second < p.`val`) {
                val (index, _) = stack.removeLast()
                res.add(index to p.`val`)
            }
            stack.add(i to p.`val`)
            i++
            p = p.next
        }
        while (stack.isNotEmpty()) {
            val (index, _) = stack.removeLast()
            res.add(index to 0)
        }

        val ans = IntArray(res.size)
        res.forEach { (index, nextGreater) -> ans[index] = nextGreater }
        return ans
    }
}