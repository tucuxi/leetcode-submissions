class Solution {
    fun minOperations(nums: IntArray, k: Int): Int {
        val s = nums.toSet()
        return when {
            s.min() < k -> -1
            k in s -> s.size - 1
            else -> s.size
        }
    }
}