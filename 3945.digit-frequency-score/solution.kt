class Solution {
    fun digitFrequencyScore(n: Int): Int {
        var res = 0
        var r = n

        while (r > 0) {
            res += r % 10
            r /= 10
        }

        return res
    }
}