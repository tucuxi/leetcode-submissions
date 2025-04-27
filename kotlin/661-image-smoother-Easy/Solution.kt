class Solution {
    fun imageSmoother(img: Array<IntArray>): Array<IntArray> {
        val m = img.size
        val n = img[0].size
        val res = Array(m) { IntArray(n) }

        fun avg(row: Int, col: Int): Int {
            var sum = 0
            var count = 0
            for (r in maxOf(0, row - 1) .. minOf(m - 1, row + 1)) {
                for (c in maxOf(0, col - 1) .. minOf(n - 1, col + 1)) {
                    sum += img[r][c]
                    count++
                }
            }
            return sum / count
        }

        for (row in 0 until m) {
            for (col in 0 until n) {
                res[row][col] = avg(row, col)
            }
        }
        
        return res
    }
}