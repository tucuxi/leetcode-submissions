class Solution {
    fun countMajoritySubarrays(nums: IntArray, target: Int): Int {
        var res = 0
        for (i in nums.indices) {
            var count = 0
            for (j in i .. nums.lastIndex) {
                if (nums[j] == target) {
                    count++
                }
                if (count * 2 > j - i + 1) {
                    res++
                }
            }
        }
        return res
    }
}