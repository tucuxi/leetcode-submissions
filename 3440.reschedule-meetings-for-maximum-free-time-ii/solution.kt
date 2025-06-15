class Solution {
    fun maxFreeTime(eventTime: Int, startTime: IntArray, endTime: IntArray): Int {
        val n = endTime.size

        val gaps = mutableListOf(0 to startTime[0])
        (1 until n).mapTo(gaps) { i -> i to startTime[i] - endTime[i - 1] }
        gaps.add(n to eventTime - endTime[n - 1])

        val orderedGaps = gaps.sortedBy { (_, duration) -> duration }

        val freeTimes = sequence {

            yield(orderedGaps[orderedGaps.lastIndex].second)
            
            (0 until n).forEach { i ->
            
                val meetingDuration = endTime[i] - startTime[i]
                val adjacentGaps = gaps[i].second + gaps[i + 1].second

                yield(adjacentGaps)

                val j = orderedGaps.binarySearchBy(meetingDuration) { (_, duration) -> duration }
                val pos = if (j >= 0) j else -(j + 1)

                if ((pos .. min(n, pos + 2)).any { k -> orderedGaps[k].first !in i .. i + 1 }) {
                    yield(meetingDuration + adjacentGaps)
                }
            }
        }

        return freeTimes.max()
    }
}