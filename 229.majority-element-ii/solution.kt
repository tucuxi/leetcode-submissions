class Solution {
    fun majorityElement(nums: IntArray): List<Int> {
        var candidate1 = 0
        var candidate2 = 1
        var count1 = 0
        var count2 = 0
        
        for (num in nums) {
            if (num == candidate1) {
                count1++
            } else if (num == candidate2) {
                count2++
            } else if (count1 == 0) {
                candidate1 = num
                count1 = 1
            } else if (count2 == 0) {
                candidate2 = num
                count2 = 1
            } else {
                count1--
                count2--
            }
        }

        return listOf(candidate1, candidate2).filter { num ->
            nums.count { it == num } > nums.size / 3
        }
    }
}