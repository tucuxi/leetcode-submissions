const val DELIM = " "

class Solution {
    fun reverseWords(s: String): String {
        val words = s.split(DELIM)
        val vowels = countVowels(words[0])
        val mappedWords = words.mapIndexed { i, s ->
            if (i == 0 || countVowels(s) != vowels) s else s.reversed()
        }
        return mappedWords.joinToString(DELIM)
    }

    fun countVowels(s: String): Int = s.count { it in "aeiou" }
}

