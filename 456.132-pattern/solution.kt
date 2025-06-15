class Solution {
    fun find132pattern(nums: IntArray): Boolean {
        var nj1 = Int.MAX_VALUE
        var nj2 = Int.MIN_VALUE
        var ni = nums[0]
        for (k in 1..nums.lastIndex) {
            val nk = nums[k]
            if (nk in nj1..nj2) {
                return true
            }
            if (ni + 1 <= nk - 1) {
                nj1 = minOf(nj1, ni + 1)
                nj2 = maxOf(nj2, nk - 1)
            }
            ni = minOf(ni, nk)
        }
        return false
    }
}