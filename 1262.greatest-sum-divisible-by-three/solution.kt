class Solution {
    fun maxSumDivThree(nums: IntArray): Int {
        var dp = IntArray(3)
        nums.forEach { n ->
            val dpnext = dp.copyOf()
            repeat(3) { i ->
                val v = dp[i] + n
                val k = v % 3
                dpnext[k] = maxOf(dpnext[k], v) 
            }
            dp = dpnext
        }
        return dp[0]
    }
}