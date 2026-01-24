class Solution {
    fun findRightInterval(intervals: Array<IntArray>): IntArray {
        val startToIndex = TreeMap<Int, Int>().also { map ->
            intervals.forEachIndexed { i, (start) -> map[start] = i }
        }
        return IntArray(intervals.size) { i ->
            startToIndex.ceilingEntry(intervals[i][1])?.value ?: -1
        }
    }
}