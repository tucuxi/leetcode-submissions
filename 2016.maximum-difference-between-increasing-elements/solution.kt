class Solution {
    fun maximumDifference(nums: IntArray): Int {
        val (_, maxDiff) = nums.fold(nums.first() to 0) { (minNum, maxDiff), num ->
            minOf(minNum, num) to maxOf(maxDiff, num - minNum)
        }
        return if (maxDiff == 0) -1 else maxDiff
    }
}