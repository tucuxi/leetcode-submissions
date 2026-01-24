import java.math.BigInteger

class Solution {
    fun minAllOneMultiple(k: Int): Int {
        if (k % 2 == 0 || k % 5 == 0) {
            return -1
        }
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