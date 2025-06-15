class Solution {
    fun runningSum(nums: IntArray): IntArray {
        var acc = 0
        return nums.map {
            acc += it
            acc
        }.toIntArray()
    }
}