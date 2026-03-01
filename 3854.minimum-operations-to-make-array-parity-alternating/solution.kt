class Solution {
    fun makeParityAlternating(nums: IntArray): IntArray {
        val n = nums.size
        if (n <= 1) {
            return intArrayOf(0, 0)
        }

        var cost1 = 0
        for (i in 0 until n) {
            val numParity = (nums[i] % 2 + 2) % 2
            val targetParity = i % 2
            if (numParity != targetParity) {
                cost1++
            }
        }

        val cost2 = n - cost1
        val minOps = minOf(cost1, cost2)
        var minRange = Long.MAX_VALUE

        if (cost1 == minOps) {
            minRange = minOf(minRange, calculateMinRange(nums, 0))
        }
        if (cost2 == minOps) {
            minRange = minOf(minRange, calculateMinRange(nums, 1))
        }

        return intArrayOf(minOps, minRange.toInt())
    }

        private fun calculateMinRange(nums: IntArray, pattern: Int): Long {
        val n = nums.size
        val allOptions = mutableListOf<Pair<Long, Int>>()

        for (i in 0 until n) {
            val numParity = (nums[i] % 2 + 2) % 2
            val targetParity = (i + pattern) % 2

            if (numParity == targetParity) {
                allOptions.add(nums[i].toLong() to i)
            } else {
                allOptions.add(nums[i].toLong() - 1 to i)
                allOptions.add(nums[i].toLong() + 1 to i)
            }
        }

        allOptions.sortBy { it.first }

        val counts = mutableMapOf<Int, Int>()
        var left = 0
        var minRange = Long.MAX_VALUE

        for (right in allOptions.indices) {
            val (_, idxR) = allOptions[right]
            counts[idxR] = counts.getOrDefault(idxR, 0) + 1

            while (counts.size == n) {
                val (valR, _) = allOptions[right]
                val (valL, idxL) = allOptions[left]
                minRange = minOf(minRange, valR - valL)

                counts[idxL] = counts[idxL]!! - 1
                if (counts[idxL] == 0) {
                    counts.remove(idxL)
                }
                left++
            }
        }
        return minRange
    }
}