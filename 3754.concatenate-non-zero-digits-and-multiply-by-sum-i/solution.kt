class Solution {
    fun sumAndMultiply(n: Int): Long {
        var sum = 0
        var x = 0L
        var p = 1
        var n = n

        while (n > 0) {
            val d = n % 10
            if (d != 0) {
                x += p * d
                p *= 10
                sum += d
            }
            n /= 10
        }

        return x * sum
    }
}