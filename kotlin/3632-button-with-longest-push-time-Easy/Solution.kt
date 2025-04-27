class Solution {
    fun buttonWithLongestTime(events: Array<IntArray>): Int {
        var longestButtonIndex = -1
        var longestButtonTime = 0
        var previousButtonTime = 0

        events.forEach { (index, time) ->
            val thisButtonTime = time - previousButtonTime
            if (thisButtonTime > longestButtonTime || thisButtonTime == longestButtonTime && index < longestButtonIndex) {
                longestButtonIndex = index
                longestButtonTime = thisButtonTime
            }
            previousButtonTime = time
        }

        return longestButtonIndex
    }
}