class Solution {
    fun minimumAbsDifference(arr: IntArray): List<List<Int>> {
        val sorted = arr.sorted()
        val windowed = sorted.windowed(2)
        val min = windowed.map { it[1] - it[0] }.min()
        return windowed.filter { it[1] - it[0] == min }.map { listOf(it[0], it[1]) }
    }
}