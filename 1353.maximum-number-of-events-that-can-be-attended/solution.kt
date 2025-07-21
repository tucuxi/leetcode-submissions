class Solution {
    fun maxEvents(events: Array<IntArray>): Int {
        val orderedEvents = events.sortedBy { it[0] }
        val lastDay = events.maxOf { it[1] }
        val pq = PriorityQueue<Int>()
        var res = 0
        var i = 0
        for (day in 1..lastDay) {
            while (i < orderedEvents.size && orderedEvents[i][0] <= day) {
                pq.offer(orderedEvents[i][1])
                i++
            }
            while (pq.isNotEmpty() && pq.peek() < day) {
                pq.poll()
            }
            if (pq.isNotEmpty()) {
                pq.poll()
                res++
            }
        }
        return res
    }
}