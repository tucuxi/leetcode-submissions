class Solution {
    fun maxSideLength(mat: Array<IntArray>, threshold: Int): Int {
        val m = mat.size
        val n = mat[0].size
        val prefix = Array(m + 1) { IntArray(n + 1) }

        for (i in 1..m) {
            for (j in 1..n) {
                prefix[i][j] = mat[i - 1][j - 1] + 
                               prefix[i - 1][j] + 
                               prefix[i][j - 1] - 
                               prefix[i - 1][j - 1]
            }
        }

        var maxSide = 0

        for (i in 1..m) {
            for (j in 1..n) {
                val target = maxSide + 1

                if (i >= target && j >= target) {
                    val sum = prefix[i][j] - 
                              prefix[i - target][j] - 
                              prefix[i][j - target] + 
                              prefix[i - target][j - target]
                    
                    if (sum <= threshold) {
                        maxSide = target
                    }
                }
            }
        }
        
        return maxSide
    }
}