class Solution {
    fun findRelativeRanks(score: IntArray): Array<String> {
        val ordered = score.withIndex().sortedByDescending { it.value }
        val res = Array(score.size) { "" }
        ordered.forEachIndexed { rank, indexedValue ->
            res[indexedValue.index] = when (rank) {
                0 -> "Gold Medal"
                1 -> "Silver Medal"
                2 -> "Bronze Medal"
                else -> "${rank + 1}"
            }
        }
        return res
    }
}