class Solution {
    fun containsNearbyDuplicate(nums: IntArray, k: Int): Boolean {
        for (i in nums.indices) {
            for (j in 1 .. k) {
                if (i + j == nums.size) {
                    break
                }
                if (nums[i] == nums[i + j]) {
                    return true
                }
            }
        }
        return false
    }
}