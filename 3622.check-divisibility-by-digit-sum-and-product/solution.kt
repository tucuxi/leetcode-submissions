class Solution {
    fun checkDivisibility(n: Int): Boolean {
        var sum = 0
        var prod = 1
        var num = n

        while (num > 0) {
            val digit = num % 10
            sum += digit
            prod *= digit
            num /= 10
        }

        return n % (sum + prod) == 0
    }
}