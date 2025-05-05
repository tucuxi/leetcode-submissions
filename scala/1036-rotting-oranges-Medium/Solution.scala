object Solution {
    def orangesRotting(grid: Array[Array[Int]]): Int = {
        val maxRow = grid.length - 1
        val maxCol = grid(0).length -1
        
        def hasFresh: Boolean = {
            for (row <- grid) if (row.contains(1)) return true
            false
        }
        
        def hasRottenNeighbor(row: Int, col: Int): Boolean = {
            (row > 0 && grid(row - 1)(col) == 2) ||
            (row < maxRow && grid(row + 1)(col) == 2) ||
            (col > 0 && grid(row)(col - 1) == 2) ||
            (col < maxCol && grid(row)(col + 1) == 2)
        }
        
        def tick: Boolean = {
            val rotting = for {
                row <- 0 to maxRow
                col <- 0 to maxCol
                if grid(row)(col) == 1
                if hasRottenNeighbor(row, col)
            } yield (row, col)
            for ((row, col) <- rotting)
                grid(row)(col) = 2
            !rotting.isEmpty
        }
        
        var minute = 0
        while (hasFresh) {
            minute += 1
            if (!tick) return -1
        }
        minute
    }
}