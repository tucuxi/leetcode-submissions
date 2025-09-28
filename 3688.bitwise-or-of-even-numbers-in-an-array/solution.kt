class Solution {
    fun evenNumberBitwiseORs(nums: IntArray): Int {
        return nums.fold(0) { acc, num ->
            if (num % 2 == 0) acc or num else acc
        }
    }
}