class Solution {
    fun longestBalanced(nums: IntArray): Int {
        var res = 0

        for (i in 0..nums.lastIndex) {
            val v = BooleanArray(100001)
            var even = 0
            var odd = 0
                
            for (j in i..nums.lastIndex) {
                val n = nums[j]

                if (!v[n]) {
                    v[n] = true
                    if (n % 2 == 0) {
                        even++
                    } else {
                        odd++
                    }
                }
                if (even == odd) {
                    res = maxOf(res, j - i + 1)
                }
            }
        }
        return res
    }
}