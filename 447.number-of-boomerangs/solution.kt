class Solution {
    fun numberOfBoomerangs(points: Array<IntArray>): Int {
        var res = 0

        for (i in points.indices) {
            val h = mutableMapOf<Int, Int>()
            for (j in points.indices) {
                if (j == i) continue
                val d = distance(points[i], points[j])
                h[d] = h.getOrDefault(d, 0) + 1
            }
            res += h.values.sumOf { it * (it - 1) } 
        }

        return res
    }

    inline fun distance(a: IntArray, b: IntArray): Int {
        val x = a[0] - b[0]
        val y = a[1] - b[1]
        return x * x + y * y
    }
}