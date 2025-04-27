class Solution {
    fun minGroups(intervals: Array<IntArray>): Int {
        return intervals
            .flatMap { (left, right) -> listOf(left to 1, right + 1 to -1) }
            .sortedWith(compareBy( { it.first }, { it.second })) 
            .runningFold(0) { acc, (_, change) -> acc + change }
            .max()
    }
}
