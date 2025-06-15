class Solution {
    fun maxFreqSum(s: String): Int {
        return s
            .partition { "aeiou".indexOf(it) >= 0 }
            .toList()
            .sumOf { l -> l.groupingBy { it }.eachCount().maxOfOrNull { it.value } ?: 0 }
    }
}