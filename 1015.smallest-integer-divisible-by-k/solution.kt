class Solution {
    fun smallestRepunitDivByK(k: Int): Int {
        var r = 0
        for (i in 1..k) {
            r = (10 * r + 1) % k
            if (r == 0) {
                return i
            }
        }
        return -1
    }
}