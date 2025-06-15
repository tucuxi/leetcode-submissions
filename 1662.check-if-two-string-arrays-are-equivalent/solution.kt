class Solution {
    fun arrayStringsAreEqual(word1: Array<String>, word2: Array<String>): Boolean {
        return concat(word1).equals(concat(word2))
    }
    
    fun concat(word: Array<String>): String {
        val sb = StringBuilder()
        for (w in word) sb.append(w)
        return sb.toString()
    }
}