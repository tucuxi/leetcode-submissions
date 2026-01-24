class Solution {
    fun sumFourDivisors(nums: IntArray): Int {
        return nums.sumOf { num ->
            val divisors = mutableSetOf<Int>()
            var x = 1

            while (divisors.size <= 4 && x * x <= num) {
                if (num % x == 0) {
                    divisors.add(x)
                    divisors.add(num / x)
                }
                x++
            }

            if (divisors.size == 4) divisors.sum() else 0
        }
    }
}