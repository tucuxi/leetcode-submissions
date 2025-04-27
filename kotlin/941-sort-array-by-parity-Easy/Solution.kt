class Solution {
    fun sortArrayByParity(nums: IntArray): IntArray {
        var p = 0
        var q = nums.lastIndex
        while (p < q) {
            when {
                nums[p] % 2 == 0 -> p++
                nums[q] % 2 == 1 -> q--
                else -> nums[p] = nums[q].also { nums[q] = nums[p] }
            }
        }
        return nums
    }
}