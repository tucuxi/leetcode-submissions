class Solution {
    fun minOperations(nums: IntArray): Int {
        var ones = 0
        var g = 0

        nums.forEach { num ->
            if (num == 1) ones++
            g = gcd(g, num)
        }
        if (ones > 0) {
            return nums.size - ones
        }
        if (g > 1) {
            return -1
        }

        var minLen = nums.size
        
        for (i in nums.indices) {
            var g = 0
            for (j in i until nums.size) {
                g = gcd(g, nums[j])
                if (g == 1) {
                    minLen = minOf(minLen, j - i + 1)
                    break
                }
            }
        }

        return minLen + nums.size - 2
    }

    fun gcd(a: Int, b: Int): Int {
        var a = a
        var b = b
        while (b != 0) {
            a = b.also { b = a % b}
        }
        return a
    }
}