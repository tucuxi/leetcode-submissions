class Solution {
    fun createTargetArray(nums: IntArray, index: IntArray): IntArray {
        val res: MutableList<Int> = mutableListOf()
        for (i in nums.indices)
            res.add(index[i], nums[i])
        return res.toIntArray()
    }
}