class Solution {
    fun prefixesDivBy5(nums: IntArray): List<Boolean> {
        return nums
            .scan(0) { acc, num -> (2 * acc + num) % 5 }
            .drop(1)
            .map { it == 0 }
    }
}