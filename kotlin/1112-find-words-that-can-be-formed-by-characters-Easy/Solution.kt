class Solution {
    fun countCharacters(words: Array<String>, chars: String): Int {
        val freq = ('a'..'z').associate { ch -> ch to chars.count { it == ch } }
        
        fun good(word: String) =
            freq.all { (ch, num) -> num >= word.count { it == ch } }
 
        return words
            .filter { good(it) }
            .map { it.length }
            .sum()
    }
}