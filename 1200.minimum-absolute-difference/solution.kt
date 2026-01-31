class Solution {
    fun minimumAbsDifference(arr: IntArray): List<List<Int>> {
        val windowed = arr.sorted().windowed(2)
        val min = windowed.map { it[1] - it[0] }.min()
        return windowed.filter { it[1] - it[0] == min }.map { listOf(it[0], it[1]) }
    }
}