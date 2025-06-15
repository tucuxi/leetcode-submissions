class Solution {
    fun removeBoxes(boxes: IntArray): Int {
        val n = boxes.size
        if (n == 0) return 0
        val dp = Array(n) { Array(n) { IntArray(n + 1) } }
        for (i in 0 until n) {
            for (k in 0..i) {
                dp[i][i][k] = sq(k + 1)
            }
        }
        for (l in 2..n) {
            for (j in l - 1 until n) {
                val i = j - l + 1
                for (k in 0..i) {  
                    var res = sq(k + 1) + dp[i + 1][j][0]
                    for (m in i + 1 .. j) {
                        if (boxes[m] == boxes[i]) {
                            res = maxOf(res, dp[i + 1][m - 1][0] + dp[m][j][k + 1])
                        }
                    }
                    dp[i][j][k] = res
                }
            }
        }
        return dp[0][n - 1][0]
    }
}

inline fun sq(a: Int) = a * a