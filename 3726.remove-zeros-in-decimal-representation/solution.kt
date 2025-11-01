class Solution {
    fun removeZeros(n: Long): Long {
        var rest = n
        var factor = 1L
        var res = 0L

        while (rest > 0) {
            val digit = rest % 10
            if (digit > 0) {
                res += digit * factor
                factor *= 10
            }
            rest /= 10
        }

        return res
    }
}