class Solution {
    fun findMissingElements(nums: IntArray): List<Int> {
        nums.distinct()
        nums.sort()

        val res = mutableListOf<Int>()
        var expected = nums.first()
        
        for (n in nums) {
            res += expected until n
            expected = n + 1
        }

        return res
    }
}