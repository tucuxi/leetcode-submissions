class Solution {
    fun smallestRangeI(nums: IntArray, k: Int): Int {
        var a = nums[0] 
        var b = nums[0]
        nums.forEach {
            a = minOf(a, it)
            b = maxOf(b, it)
        }
        return maxOf(0, b - a - k - k)
    }
}