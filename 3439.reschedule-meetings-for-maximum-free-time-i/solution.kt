class Solution {
    fun maxFreeTime(eventTime: Int, k: Int, startTime: IntArray, endTime: IntArray): Int {
        val n = endTime.size

        fun gap(i: Int): Int = when {
            i < 0 -> 0
            i == 0 -> startTime[0]
            i == n -> eventTime - endTime[n - 1]
            else -> startTime[i] - endTime[i - 1]
        }

        val (_, max) = (0 .. n).fold(0 to 0) { (sum, max), i ->
            val newSum = sum + gap(i) - gap(i - k - 1)
            val newMax = maxOf(max, newSum)
            newSum to newMax
        }

        return max
    }
}