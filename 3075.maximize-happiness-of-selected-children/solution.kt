class Solution {
    fun maximumHappinessSum(happiness: IntArray, k: Int): Long {
        happiness.sortDescending()
        val selected = happiness.take(k)
        return selected.withIndex().sumOf { (i, h) ->
            maxOf(0L, (h - i).toLong())
        }
    }
}