class Solution {
    fun findGCD(nums: IntArray): Int {
        val (min, max) = findMinMax(nums)
        return gcd(min, max)
    }
    
    fun findMinMax(nums: IntArray): Pair<Int, Int> {
        var min = nums[0]
        var max = nums[0]
        for (n in nums) {
            min = minOf(min, n)
            max = maxOf(max, n)
        }
        return Pair(min, max)
    }
    
    fun gcd(num1: Int, num2: Int): Int {
        var a = num1
        var b = num2
        while (b != 0) {
            val t = b
            b = a % b
            a = t
        }
        return a
    }
}