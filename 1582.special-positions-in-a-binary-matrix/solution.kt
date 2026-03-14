class Solution {
    fun numSpecial(mat: Array<IntArray>): Int {
        val m = mat.size
        val n = mat[0].size
        val rs = IntArray(m)
        val cs = IntArray(n)
        
        for (i in 0 until m) {
            for (j in 0 until n) {
                rs[i] += mat[i][j]
                cs[j] += mat[i][j]
            }
        }

        var res = 0

        for (i in 0 until m) {
            if (rs[i] != 1) continue
            for (j in 0 until n) {
                if (cs[j] == 1 && mat[i][j] == 1) {
                    res++
                }
            }
        }

        return res
    }
}