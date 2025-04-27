import kotlin.math.abs

class Solution {
    fun countKDifference(nums: IntArray, k: Int): Int {
        var res = 0
        for (i in nums.indices) {
            for (j in i+1 .. nums.lastIndex) {
                if (abs(nums[i] - nums[j]) == k) {
                    res++
                }
            }
        }
        return res
    }
}