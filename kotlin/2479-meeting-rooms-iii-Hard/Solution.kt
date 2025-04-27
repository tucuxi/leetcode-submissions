class Solution {
    fun mostBooked(n: Int, meetings: Array<IntArray>): Int {
        val roomUseCounts = IntArray(n)
        val roomBookedUntil = IntArray(n)
        val availableRooms = PriorityQueue<Int>((0 until n).toList())
        
        val bookedRooms = PriorityQueue<Int> { r1, r2 ->
            if (roomBookedUntil[r1] == roomBookedUntil[r2]) {
                r1 - r2
            } else {
                roomBookedUntil[r1] - roomBookedUntil[r2]
            }
        }
        
        for ((start, end) in meetings.sortedBy { (start, _) -> start }) {
            while (bookedRooms.isNotEmpty() && roomBookedUntil[bookedRooms.peek()] <= start) {
                val r = bookedRooms.remove()
                availableRooms.add(r)
            }
            val (r, delay) = if (availableRooms.isEmpty()) {
                val r = bookedRooms.remove()
                Pair(r, roomBookedUntil[r] - start)
            } else {
                val r = availableRooms.remove()
                Pair(r, 0)
            }
            roomBookedUntil[r] = end + delay
            bookedRooms.add(r)
            roomUseCounts[r]++
        }

        return roomUseCounts.withIndex().maxBy { it.value }.index
    }
}