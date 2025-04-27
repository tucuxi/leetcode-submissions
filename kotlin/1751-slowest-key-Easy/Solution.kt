class Solution {
    fun slowestKey(releaseTimes: IntArray, keysPressed: String): Char {
        var maxKey = keysPressed[0]
        var maxTime = releaseTimes[0]
        for (i in 1 until releaseTimes.size) {
            val time = releaseTimes[i] - releaseTimes[i - 1]
            if (time > maxTime || time == maxTime && keysPressed[i] > maxKey) {
                maxKey = keysPressed[i]
                maxTime = time
            }
        }
        return maxKey
    }
}