class Solution {
    fun largestAltitude(gain: IntArray): Int {
        var height = 0
        var maxHeight = 0
        gain.forEach {
            height += it
            maxHeight = maxOf(height, maxHeight)
        }
        return maxHeight
    }
}