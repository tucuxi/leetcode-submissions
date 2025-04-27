class Solution {
    fun findMissingAndRepeatedValues(grid: Array<IntArray>): IntArray {
        val n = grid.size
        val h = IntArray(n * n + 1)
        grid.forEach { row ->
            row.forEach { elem ->
                h[elem]++
            }
        }
        
        val res = IntArray(2)
        h.forEachIndexed { index, value ->
            when (value) {
                0 -> res[1] = index
                2 -> res[0] = index
            }
        }

        return res
    }
}