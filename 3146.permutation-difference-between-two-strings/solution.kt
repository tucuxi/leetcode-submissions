class Solution {
    fun findPermutationDifference(s: String, t: String): Int {
        val tIndex = t.withIndex().associate { it.value to it.index }
        return s.withIndex().sumBy { abs(it.index - requireNotNull(tIndex[it.value])) }
    }
}