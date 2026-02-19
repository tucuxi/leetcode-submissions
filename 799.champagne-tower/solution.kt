class Solution {
    fun champagneTower(poured: Int, query_row: Int, query_glass: Int): Double {
        var currentRow = DoubleArray(1) { poured.toDouble() }
        repeat(query_row) {
            val nextRow = DoubleArray(currentRow.size + 1)
            for (c in 0 until nextRow.lastIndex) {
                val q = (currentRow[c] - 1) / 2
                if (q > 0) {
                    nextRow[c] += q
                    nextRow[c + 1] += q
                }
            }
            currentRow = nextRow
        }
        return minOf(1.0, currentRow[query_glass])
    }
}