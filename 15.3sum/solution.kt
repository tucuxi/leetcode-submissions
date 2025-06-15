class Solution {
    fun threeSum(nums: IntArray): List<List<Int>> {
        nums.sort()
        val res: MutableSet<List<Int>> = mutableSetOf()
        for (i in nums.indices) {
            var j = i + 1
            var k = nums.lastIndex
            while (j < k) {
                val sum = nums[i] + nums[j] + nums[k]
                if (sum < 0) {
                    j++
                } else if (sum > 0) {
                    k--
                } else {
                    res.add(listOf(nums[i], nums[j], nums[k]))
                    j++
                    k--
                }
            }
        }
        return res.toList()
    }
}