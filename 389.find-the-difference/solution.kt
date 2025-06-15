class Solution {
    fun findTheDifference(s: String, t: String): Char {
        val sfreq = s.groupingBy { it }.eachCount()
        val tfreq = t.groupingBy { it }.eachCount()
        return tfreq.entries.find { (ch, count) -> sfreq.getOrDefault(ch, 0) < count }!!.key
    }
}