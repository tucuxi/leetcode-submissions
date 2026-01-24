class Solution {
    fun largestSquareArea(bottomLeft: Array<IntArray>, topRight: Array<IntArray>): Long {
        val n = bottomLeft.size
        var maxSide = 0
        for (i in 0 until n) {
            for (j in i + 1 until n) {
                val w = minOf(topRight[i][0], topRight[j][0]) - maxOf(bottomLeft[i][0], bottomLeft[j][0])
                if (w <= 0) continue
                val h = minOf(topRight[i][1], topRight[j][1]) - maxOf(bottomLeft[i][1], bottomLeft[j][1])
                if (h <= 0) continue
                maxSide = maxOf(maxSide, minOf(w, h))
            }
        }
        return maxSide.toLong() * maxSide.toLong()
    }
}