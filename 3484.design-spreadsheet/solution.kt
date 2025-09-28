class Spreadsheet(rows: Int) {

    private val grid = Array(rows) { IntArray(26) }

    fun setCell(cell: String, value: Int) {
        val (row, col) = indexOf(cell)
        grid[row][col] = value

    }

    fun resetCell(cell: String) {
        setCell(cell, 0)
    }

    fun getValue(formula: String): Int {
        require(formula.first() == '=')
        val (x, y) = formula.drop(1).split("+")
        return evaluateTerm(x) + evaluateTerm(y) 
    }

    private fun indexOf(cell: String): Pair<Int, Int> {
        return (cell.drop(1)).toInt() - 1 to  (cell.first() - 'A')
    }

    private fun getCell(cell: String): Int {
        val (row, col) = indexOf(cell)
        return grid[row][col]
    }

    private fun evaluateTerm(term: String): Int {
        return if (term.first().isLetter()) getCell(term) else term.toInt()
    }
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * var obj = Spreadsheet(rows)
 * obj.setCell(cell,value)
 * obj.resetCell(cell)
 * var param_3 = obj.getValue(formula)
 */