class Solution {
    fun merge(intervals: Array<IntArray>): Array<IntArray> {
        val seq = mutableListOf<IntArray>()
        var lastInterval = intArrayOf(0, 0)
        for (interval in intervals.sortedBy { it[0] }) {
            if (seq.isEmpty() || interval[0] > lastInterval[1]) {
                seq += interval
                lastInterval = interval
            } else {
                lastInterval[1] = maxOf(lastInterval[1], interval[1])
            }
        } 
        return seq.toTypedArray()
    }
}