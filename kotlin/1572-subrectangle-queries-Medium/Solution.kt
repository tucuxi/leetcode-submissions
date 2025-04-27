class SubrectangleQueries(rectangle: Array<IntArray>) {
    
    private val rect = rectangle

    fun updateSubrectangle(row1: Int, col1: Int, row2: Int, col2: Int, newValue: Int) {
        for (row in row1..row2) {
            for (col in col1..col2) {
                rect[row][col] = newValue
            }
        }
    }

    fun getValue(row: Int, col: Int): Int {
        return rect[row][col]
    }

}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * var obj = SubrectangleQueries(rectangle)
 * obj.updateSubrectangle(row1,col1,row2,col2,newValue)
 * var param_2 = obj.getValue(row,col)
 */