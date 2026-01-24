class Solution {
    fun maximizeSquareHoleArea(n: Int, m: Int, hBars: IntArray, vBars: IntArray): Int {
        val gap = minOf(maxGap(hBars), maxGap(vBars)) + 1
        return gap * gap        
    }

    fun maxGap(bars: IntArray): Int {
        return if (bars.size == 1) {
            1
        } else {
            var run = 1
            bars.sorted()
                .zipWithNext()
                .maxOf { (a, b) ->
                    if (a + 1 == b) {
                        run++
                    } else {
                        run = 1
                    }
                    run
                }
        }
    }
}