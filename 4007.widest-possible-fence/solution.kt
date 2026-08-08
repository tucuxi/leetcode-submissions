class Solution {
    fun maximumWidth(planks: IntArray): Int {
        val frequencies = planks.asIterable().groupingBy { it }.eachCount()
        val widthForHeight = frequencies.toMutableMap()
        val uniquePlanks = frequencies.keys.toList()

        uniquePlanks.forEachIndexed { i, p1 ->
            val frequency1 = frequencies.getOrDefault(p1, 0)
            if (frequency1 >= 2) {
                val frequency2 = widthForHeight.getOrDefault(2 * p1, 0)
                widthForHeight[2 * p1] = frequency2 + frequency1 / 2
            }
            for (j in i + 1 until uniquePlanks.size) {
                val p2 = uniquePlanks[j]
                val sum = p1 + p2
                val frequency2 = widthForHeight.getOrDefault(sum, 0)
                widthForHeight[sum] = frequency2 + minOf(frequency1, frequencies.getValue(p2))
            }
        }

        return widthForHeight.values.max()
    }
}