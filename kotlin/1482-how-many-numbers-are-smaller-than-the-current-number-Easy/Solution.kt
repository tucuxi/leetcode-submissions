class Solution {
    fun smallerNumbersThanCurrent(nums: IntArray): IntArray {
        val res = IntArray(nums.size)
        for (i in res.indices) {
            for (j in nums.indices) {
                if (nums[j] < nums[i])
                    res[i] += 1
            }            
        }
        return res
    }
}