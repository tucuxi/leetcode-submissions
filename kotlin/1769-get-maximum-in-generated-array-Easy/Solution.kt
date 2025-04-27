class Solution {
    fun getMaximumGenerated(n: Int): Int {
        if (n < 2) return n
        val nums = IntArray(n + 1)
        nums[0] = 0
        nums[1] = 1
        var res = 1
        for (i in 2..n) {
            nums[i] = if (i % 2 == 0) nums[i / 2] else nums[i / 2] + nums[i / 2 + 1]
            if (nums[i] > res) res = nums[i]
        }
        return res
    }
}