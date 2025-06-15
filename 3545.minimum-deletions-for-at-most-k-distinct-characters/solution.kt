class Solution {
    fun minDeletion(s: String, k: Int): Int {
        val c = s
            .groupingBy { it }
            .eachCount()
            .values
            .sorted()
        return c.take(maxOf(0, c.size - k)).sum()
    }
}