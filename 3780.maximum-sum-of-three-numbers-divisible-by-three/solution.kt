class Solution {
    fun maximumSum(nums: IntArray): Int {
        val g = Array(3) { PriorityQueue<Int>(4) }

        for (num in nums) {
            val i = num % 3
            g[i].offer(num)
            if (g[i].size > 3) {
                g[i].poll()
            }
        }

        val h = g.map { it.toIntArray() }

        var res = 0

        for (r in h) {
            if (r.size == 3) {
                res = maxOf(res, r.sum())
            }
        }

        if (h[0].isNotEmpty() && h[1].isNotEmpty() && h[2].isNotEmpty()) {
            res = maxOf(res, h[0].max() + h[1].max() + h[2].max())
        }

        return res
    }
}