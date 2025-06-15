class Solution {
    fun rectangleArea(rectangles: Array<IntArray>): Int {

        fun makeEvents(): List<Event> {
            val events = mutableListOf<Event>()
            for (rec in rectangles) {
                events += Event(rec[1], true, rec[0], rec[2])
                events += Event(rec[3], false, rec[0], rec[2])
            }
            return events.sortedBy { it.y }
        }

        fun sweep(events: List<Event>): Long {
            val active = mutableListOf<Pair<Int, Int>>()
            var curY = events[0].y
            var res = 0L
            for (event in events) {
                var query = 0L
                var cur = -1
                for (xs in active) {
                    cur = maxOf(cur, xs.first);
                    query += maxOf(xs.second - cur, 0)
                    cur = maxOf(cur, xs.second)
                }
                res += query * (event.y - curY)
                if (event.open) {
                    active.add(event.x1 to event.x2)
                    active.sortBy { it.first }
                } else {
                    active.remove(event.x1 to event.x2)
                }
                curY = event.y
            }
            return res
        } 
        
        val events = makeEvents()
        val res = sweep(events)
        return (res % 1_000_000_007).toInt()    
    }
}

data class Event(val y: Int, val open: Boolean, val x1: Int, val x2: Int)