class Solution {
    fun matrixReshape(mat: Array<IntArray>, r: Int, c: Int): Array<IntArray> {
        if (mat.size * mat[0].size !=  r * c) return mat
        val res = Array(r) { IntArray(c) }
        var ir = 0
        var ic = 0
        for (i in mat.indices) {
            for (j in mat[i].indices) {
                res[ir][ic] = mat[i][j]
                if (++ic == c) {
                    ic = 0
                    ir++
                }
            }
        }
        return res
    }
}