class Solution {
    fun smallestIndex(nums: IntArray): Int =
        nums.withIndex()
            .firstOrNull { (index, num) -> index == digitSum(num) }
            ?.index
            ?: -1

    tailrec fun digitSum(num: Int, acc: Int = 0): Int =
        if (num == 0) acc else digitSum(num / 10, acc + num % 10)
}