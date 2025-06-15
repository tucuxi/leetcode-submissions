class Solution {
    fun findWordsContaining(words: Array<String>, x: Char): List<Int> {
        return words.indices.filter { i -> words[i].indexOf(x) != -1 }
    }
}