class Solution {
    fun countCoveredBuildings(n: Int, buildings: Array<IntArray>): Int {
        val minRow = IntArray(n + 1) { n + 1 }
        val maxRow = IntArray(n + 1)
        val minCol = IntArray(n + 1) { n + 1 }
        val maxCol = IntArray(n + 1)

        buildings.forEach { (x, y) ->
            minRow[y] = minOf(minRow[y], x)
            maxRow[y] = maxOf(maxRow[y], x)
            minCol[x] = minOf(minCol[x], y)
            maxCol[x] = maxOf(maxCol[x], y)
        }
        
        return buildings.count { (x, y) -> x > minRow[y] && x < maxRow[y] && y > minCol[x] && y < maxCol[x] }
    }
}