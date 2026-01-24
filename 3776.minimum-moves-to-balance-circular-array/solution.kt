class Solution {
    fun minMoves(balance: IntArray): Long {
        val negative = balance.withIndex().firstOrNull { (_, b) -> b < 0 } ?: return 0
        
        val distanceMap: Map<Int, Int> = balance
            .mapIndexed { i, b ->
                val d = abs(i - negative.index)
                minOf(d, balance.size - d) to b
            }
            .groupingBy { (i, _) -> i }
            .aggregate { _, acc, (_, b), _ -> (acc ?: 0) + b }
        
        val orderedDistances = distanceMap
            .entries
            .sortedBy { (key, _) -> key }
            .drop(1)
            .toMutableList()
        
        var rest = -negative.value
        var moves = 0L
        
        while (orderedDistances.isNotEmpty() && rest > 0) {
            val (distance, count) = orderedDistances.removeAt(0)
            val move = minOf(rest, count)
            rest -= move
            moves += move.toLong() * distance.toLong()
        }
        return if (rest == 0) moves else -1
    }
}