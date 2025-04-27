class Solution {
    fun isToeplitzMatrix(matrix: Array<IntArray>): Boolean {
        for (i in 1 until matrix.size)
            for (j in 1 until matrix[i].size)
                if (matrix[i][j] != matrix[i - 1][j - 1])
                    return false
        return true
    }
}