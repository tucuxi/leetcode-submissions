class Solution {
    fun findPoisonedDuration(timeSeries: IntArray, duration: Int): Int {
        var total = duration
        for (i in 1..timeSeries.lastIndex) {
            total += minOf(duration, timeSeries[i] - timeSeries[i - 1])
        }
        return total
    }
}