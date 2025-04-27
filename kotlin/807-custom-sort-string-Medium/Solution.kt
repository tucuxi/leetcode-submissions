class Solution {
    fun customSortString(order: String, s: String): String {
        val index = order.withIndex().associate { it.value to it.index }
        return s.toList().sortedBy { index[it] }.joinToString("")
    }
}