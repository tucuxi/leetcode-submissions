import kotlin.math.abs

class Solution {
    fun minTimeToType(word: String): Int {
        var seconds = 0
        var prev = 'a'
        for (ch in word) {
            val steps = abs(ch - prev) 
            seconds += minOf(steps, 26 - steps) + 1
            prev = ch
        }
        return seconds
    }
}