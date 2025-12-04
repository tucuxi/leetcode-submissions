class Solution {
    fun minSubarray(nums: IntArray, p: Int): Int {
        val remainder = nums.fold(0) { acc, num -> (acc + num) % p }

        if (remainder == 0) {
            return 0
        }

        val lastOccurrence = mutableMapOf(0 to -1)
        var currentSum = 0
        var res = nums.size

        nums.forEachIndexed { i, num ->
            currentSum = (currentSum + num) % p
            val needed = (currentSum - remainder + p) % p
            lastOccurrence[needed]?.let { j -> res = minOf(res, i - j) }
            lastOccurrence[currentSum] = i
        }

        return if (res == nums.size) -1 else res
    }
}