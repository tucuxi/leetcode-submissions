class Solution {
    fun sortColors(nums: IntArray) {
        val counts = nums.fold(arrayOf(0, 0, 0)) { acc, num -> acc.also { it[num]++ }}
        counts.withIndex().fold(0) { i, (color, count) -> 
            repeat(count) {
                nums[i + it] = color
            }
            i + count
        }
    }
}