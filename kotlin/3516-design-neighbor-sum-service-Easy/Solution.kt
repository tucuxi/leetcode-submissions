class NeighborSum(grid: Array<IntArray>) {

    private val values = grid

    private val position = grid.flatMapIndexed { i, row ->
        row.mapIndexed { j, v ->
            v to Pair(i, j)
        }
    }.toMap()

    fun adjacentSum(value: Int): Int {
        val (i, j) = requireNotNull(position[value])
        return (if (i > 0) values[i - 1][j] else 0) +
            (if (i < values.lastIndex) values[i + 1][j] else 0) +
            (if (j > 0) values[i][j - 1] else 0) +
            (if (j < values.lastIndex) values[i][j + 1] else 0)
    }

    fun diagonalSum(value: Int): Int {
        val (i, j) = requireNotNull(position[value])
        return (if (i > 0 && j > 0) values[i - 1][j - 1] else 0) +
            (if (i < values.lastIndex && j > 0) values[i + 1][j - 1] else 0) +
            (if (i > 0 && j < values.lastIndex) values[i - 1][j + 1] else 0) +
            (if (i < values.lastIndex && j < values.lastIndex) values[i + 1][j + 1] else 0)
    }

}

/**
 * Your NeighborSum object will be instantiated and called as such:
 * var obj = NeighborSum(grid)
 * var param_1 = obj.adjacentSum(value)
 * var param_2 = obj.diagonalSum(value)
 */