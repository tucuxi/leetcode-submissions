class Solution {
    fun sumDivisibleByK(nums: IntArray, k: Int): Int {
        return nums
            .asIterable()
            .groupingBy { it }
            .eachCount()
            .filterValues { it % k == 0 }
            .asIterable()
            .sumBy { it.key * it. value }
    }
}