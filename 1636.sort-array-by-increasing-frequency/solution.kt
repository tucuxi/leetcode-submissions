class Solution {
    fun frequencySort(nums: IntArray): IntArray {
        val count = IntArray(201)
        nums.forEach { count[it + 100]++ }
        return nums.sortedWith(compareBy<Int> { count[it + 100] }.thenByDescending { it }).toIntArray()
    }
}