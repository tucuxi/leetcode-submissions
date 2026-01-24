class Solution(val rects: Array<IntArray>) {

    val areas = rects
        .scan(0L) { acc, (a, b, x, y) -> acc + (x - a + 1) * (y - b + 1) }
        .drop(1)

    fun pick(): IntArray {
        val r = (1..areas.last()).random()
        val i = areas.binarySearch(r).let { if (it < 0) -(it + 1) else it }
        return rects[i].let { (a, b, x, y) ->
            intArrayOf((a..x).random(), (b..y).random())
        }
    }

}

/**
 * Your Solution object will be instantiated and called as such:
 * var obj = Solution(rects)
 * var param_1 = obj.pick()
 */