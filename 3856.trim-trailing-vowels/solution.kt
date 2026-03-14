class Solution {
    fun trimTrailingVowels(s: String): String {
        return s.dropLastWhile { "aeiou".indexOf(it) != -1 }
    }
}