class Solution {
    fun minFlipsMonoIncr(s: String): Int {
        var ones = 0
        var flips = 0
        for (ch in s) {
            if (ch == '0') {
                flips++
            } else {
                ones++
            }
            flips = minOf(flips, ones)
        }
        return flips
    }
}