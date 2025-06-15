class Solution {
    fun numIdenticalPairs(nums: IntArray): Int {
        var res = 0
        for (i in nums.indices)
            for (j in 0 until i)
                if (nums[i] == nums[j])
                    res += 1
        return res
    }
}