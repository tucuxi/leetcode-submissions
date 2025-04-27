class Solution {
    fun findShortestSubArray(nums: IntArray): Int {
        
        fun findIndex(num: Int, indices: IntProgression): Int {
            for (i in indices) {
                if (nums[i] == num) {
                    return i
                }
            }
            return -1
        }

        val frequencies = nums.groupBy { it }.map { (num, elems) -> num to elems.size }
        val descendingFrequencies = frequencies.sortedByDescending { (_, freq) -> freq }
        val mostFrequentNums = descendingFrequencies.takeWhile { (_, freq) -> freq == descendingFrequencies[0].second }.map { (num, _) -> num }
        val reversedIndices = nums.indices.reversed()
        return mostFrequentNums.map { num ->
            findIndex(num, reversedIndices) - findIndex(num, nums.indices)
        }.min()!! + 1
    }
}