class Solution {
    fun stringMatching(words: Array<String>): List<String> {
        return words.filter { w -> words.any { it != w && it.contains(w) }}
    }
}