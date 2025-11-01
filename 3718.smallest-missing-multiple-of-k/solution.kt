class Solution {
    fun missingMultiple(nums: IntArray, k: Int): Int {
        val numsSet = nums.toSet()
        return (k .. 100 + k step k).first { it !in numsSet }
    }
}