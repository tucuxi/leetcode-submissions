class Solution {
    fun percentageLetter(s: String, letter: Char): Int {
        val freq = s.count { it == letter }
        return freq * 100 / s.length
        
    }
}