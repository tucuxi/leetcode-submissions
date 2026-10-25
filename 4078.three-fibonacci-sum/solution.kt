class Solution {
    fun threeFibonacciSum(n: Int): Boolean {
        var a = 0
        var b = 1

        while (true) {
            val c = a + b
            if (c + c == n) {
                return true
            }
            if (c + c > n) {
                return false
            }
            a = b
            b = c
        }
    }
}