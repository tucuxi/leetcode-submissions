class Solution {
    fun isGood(nums: IntArray): Boolean {
        val n = nums.size - 1
        val c = IntArray(nums.size)
        for (num in nums) {
            if (num > n || num < n && c[num] > 0 || num == n && c[num] > 1) {
                return false
            }
            c[num]++
        }
        return true
    }
}