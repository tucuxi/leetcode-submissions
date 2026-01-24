class Solution {
    fun maxTwoEvents(events: Array<IntArray>): Int {
        val pq = PriorityQueue<IntArray>(compareBy { it[1] })
        var maxValue = 0
        var maxSum = 0

        events.sortBy { it[0] }
        events.forEach { event ->
            while (pq.isNotEmpty() && pq.peek()[1] < event[0] ) {
                maxValue = maxOf(maxValue, pq.poll()[2])
            }
            maxSum = maxOf(maxSum, maxValue + event[2])
            pq.offer(event)
        }

        return maxSum
    }
}