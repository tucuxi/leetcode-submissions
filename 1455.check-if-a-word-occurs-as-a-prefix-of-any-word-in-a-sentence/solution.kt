class Solution {
    fun isPrefixOfWord(sentence: String, searchWord: String): Int {
        return sentence
            .split(" ")
            .withIndex()
            .find { (_, word) -> word.startsWith(searchWord) }
            ?.let { (index, _) -> index + 1 }
            ?: -1
    }
}