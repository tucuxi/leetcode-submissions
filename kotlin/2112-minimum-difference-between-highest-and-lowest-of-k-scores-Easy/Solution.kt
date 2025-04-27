class Solution {
    fun minimumDifference(nums: IntArray, k: Int): Int {
        val sortedNums = nums.sorted()
        var min = sortedNums[k - 1] - sortedNums[0]
        for (i in k .. sortedNums.lastIndex) {
            min = minOf(min, sortedNums[i] - sortedNums[i - k + 1])
        }
        return min
    }
}