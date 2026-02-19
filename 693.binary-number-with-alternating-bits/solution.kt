class Solution {
    fun hasAlternatingBits(n: Int): Boolean {
        var p = n

        while (p > 0) {
            val q = p shr 1
            if (p and 1 == q and 1) {
                return false
            }
            p = q
        }
        return true
    }
}