class Solution {
    fun areOccurrencesEqual(s: String): Boolean {
        val freq = s.groupingBy { it }.eachCount().values
        return freq.all { it == freq.first() } 
    }
}