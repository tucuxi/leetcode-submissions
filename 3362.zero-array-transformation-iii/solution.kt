class Solution {
    fun maxRemoval(nums: IntArray, queries: Array<IntArray>): Int {
        val pq = PriorityQueue<Int>(Comparator.reverseOrder())
        val deltaArray = IntArray(nums.size + 1)
        var operations = 0
        var j = 0

        queries.sortBy { (l, _) -> l }
        nums.withIndex().forEach { (i, n) ->
            operations += deltaArray[i]
            while (j < queries.size && queries[j][0] == i) {
                pq.offer(queries[j][1])
                j++
            }
            while (operations < n && pq.size > 0 && pq.peek() >= i) {
                operations++
                deltaArray[pq.poll() + 1]--
            }
            if (operations < n) {
                return -1
            }
        }

        return pq.size
    }
}