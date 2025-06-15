class Solution {
    fun sumOfUnique(nums: IntArray): Int =
        nums.groupBy { it }
            .mapNotNull { (num, values) -> if (values.size == 1) num else null }
            .sum()
}