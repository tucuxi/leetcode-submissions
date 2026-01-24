class Solution {
    fun centeredSubarrays(nums: IntArray): Int {
        var res = 0
        for (i in nums.indices) {
            val elements = mutableSetOf<Int>()
            var sum = 0
            for (j in i until nums.size) {
                elements.add(nums[j])
                sum += nums[j]
                if (sum in elements) {
                    res++
                }
            }
        }
        return res
    }
}