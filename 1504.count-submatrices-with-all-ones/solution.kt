class Solution {
    fun numSubmat(mat: Array<IntArray>): Int {
        val n = mat[0].size
        val row = Array(mat.size) { IntArray(n) }
        var res = 0

        for (i in mat.indices) {
            for (j in 0 until n) {
                row[i][j] = if (j == 0) {
                    mat[i][j]
                } else if (mat[i][j] == 0) {
                    0
                } else {
                    row[i][j - 1] + 1
                }
                var cur = row[i][j]
                for (k in i downTo 0) {
                    cur = minOf(cur, row[k][j])
                    res += cur
                }
            }
        }
        return res
    }
}