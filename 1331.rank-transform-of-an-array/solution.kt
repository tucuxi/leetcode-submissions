class Solution {
    fun arrayRankTransform(arr: IntArray): IntArray {
        val h = arr.toSet().sorted().withIndex().associate { it.value to it.index + 1 }
        return IntArray(arr.size) { h[arr[it]]!! }
    }
}