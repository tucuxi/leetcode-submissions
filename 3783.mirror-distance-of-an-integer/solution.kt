class Solution {
    fun mirrorDistance(n: Int): Int {
        return abs(n - reverse(n))
    }

    fun reverse(n: Int): Int {
        var r = 0
        var k = n
        while (k > 0) {
            r = 10 * r + k % 10
            k /= 10
        }
        return r
    }
}