class Solution {
    fun twoSum(nums: IntArray, target: Int): IntArray {
        val map = mutableMapOf<Int, Int>()
        nums.forEachIndexed { index, num ->
            val other = map[target - num]
            if (other != null) return intArrayOf(index, other)
            map[num] = index
        }
        return IntArray(0)
    }
}