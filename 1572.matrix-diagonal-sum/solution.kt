class Solution {
    fun diagonalSum(mat: Array<IntArray>): Int {
        val n = mat.size
        val d = if (n % 2 == 0) 0 else -mat[n / 2][n / 2]
        return mat.foldIndexed(d) {
            index, acc, elem -> acc + elem[index] + elem[n - index - 1]
        }
    }
}