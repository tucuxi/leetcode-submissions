class Solution {
    fun trapRainWater(heightMap: Array<IntArray>): Int {
        val nRows = heightMap.size
        val nCols = heightMap.first().size
        val visited = Array(nRows) { BooleanArray(nCols) }
        val boundary = PriorityQueue<Cell>()

        (0 until nRows).forEach { i ->
            boundary.offer(Cell(heightMap[i][0], i, 0))
            visited[i][0] = true
            boundary.offer(Cell(heightMap[i][nCols - 1], i, nCols - 1))
            visited[i][nCols - 1] = true
        }
        (0 until nCols).forEach { j->
            boundary.offer(Cell(heightMap[0][j], 0, j))
            visited[0][j] = true
            boundary.offer(Cell(heightMap[nRows - 1][j], nRows - 1, j))
            visited[nRows - 1][j] = true
        }

        var totalWaterVolume = 0

        while (boundary.isNotEmpty()) {
            val currentCell = boundary.poll()
            directions.forEach { (dRow, dCol) ->
                val neighborRow = currentCell.row + dRow
                val neighborCol = currentCell.col + dCol
                if (neighborRow in 0 until nRows && neighborCol in 0 until nCols && !visited[neighborRow][neighborCol]) {
                    val neighborHeight = heightMap[neighborRow][neighborCol]
                    totalWaterVolume += maxOf(0, currentCell.height - neighborHeight)
                    boundary.offer(Cell(maxOf(neighborHeight, currentCell.height), neighborRow, neighborCol))
                    visited[neighborRow][neighborCol] = true
                }
            }
        }

        return totalWaterVolume
    }

    data class Cell(val height: Int, val row: Int, val col: Int) : Comparable<Cell> {
        override fun compareTo(other: Cell): Int = height.compareTo(other.height)
    }

    companion object {
        val directions = arrayOf(0 to -1, 0 to 1, -1 to 0, 1 to 0)
    }
}
