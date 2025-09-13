class Solution {
    fun doesAliceWin(s: String): Boolean {
        return s.any { "aeiou".indexOf(it) != -1 }
    }
}