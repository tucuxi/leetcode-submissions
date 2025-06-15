class Solution {
    fun frequencySort(s: String): String {
        return s
            .groupingBy { it }
            .eachCount()
            .entries
            .sortedByDescending { (_, count) -> count }
            .map { (ch, count) -> ch.toString().repeat(count) }
            .joinToString("")
    }
}