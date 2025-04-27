class Solution {
    fun prefixCount(words: Array<String>, pref: String): Int =
        words.count { it.startsWith(pref) }
}