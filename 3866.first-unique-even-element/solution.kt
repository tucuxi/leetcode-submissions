class Solution {
    fun firstUniqueEven(nums: IntArray): Int {
        val counts = IntArray(101)
        nums.forEach { counts[it]++ }
        return nums.firstOrNull { it % 2  == 0 && counts[it] == 1 } ?: -1
    }
}