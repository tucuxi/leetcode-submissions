class Solution {
    fun oddCells(n: Int, m: Int, indices: Array<IntArray>): Int {
        val mat = Array(n) { IntArray(m) }
        for (index in indices) {
            val ri = index[0]
            for (col in 0 until m)
                mat[ri][col]++
            val ci = index[1]
            for (row in 0 until n)
                mat[row][ci]++
        }
        return mat.fold(0) { sum, row -> sum + row.count { it % 2 == 1 } }
    }
}