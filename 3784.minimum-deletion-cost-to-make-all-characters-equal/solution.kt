class Solution {
    fun minCost(s: String, cost: IntArray): Long {
        var totalSum = 0L
        val charSums = LongArray(26)

        s.indices.forEach { i ->
            val charIndex = s[i] - 'a'
            val currentCost = cost[i].toLong()
            totalSum += currentCost
            charSums[charIndex] += currentCost
        }

        val maxCharSum = charSums.maxOrNull() ?: 0L
        return totalSum - maxCharSum
    }
}