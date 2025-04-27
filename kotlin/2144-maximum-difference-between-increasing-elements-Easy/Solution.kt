class Solution {
    fun maximumDifference(nums: IntArray): Int {
        var min = nums[0]
        var res = -1
        for (n in nums) {
            if (min < n) {
                res = maxOf(res, n - min)
            } else {
                min = n
            }
        }
        return res
    }
}