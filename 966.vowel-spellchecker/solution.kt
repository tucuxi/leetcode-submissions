class Solution {
    fun spellchecker(wordlist: Array<String>, queries: Array<String>): Array<String> {
        return with(SpellChecker(wordlist)) { queries.map { find(it) } }.toTypedArray()
    }
}

class SpellChecker(wordlist: Array<String>) {
    private val byLiteral = wordlist.toSet()
    private val byLowerCase = mutableMapOf<String, String>()
    private val byDevowelized = mutableMapOf<String, String>()

    init {
        wordlist.forEach { s->
            byLowerCase.putIfAbsent(lowerCase(s), s)
            byDevowelized.putIfAbsent(devowelized(s), s)
        }
    }

    fun find(word: String): String {
        return if (byLiteral.contains(word)) {
            word
        } else {
            byLowerCase.getOrElse(lowerCase(word)) {
                byDevowelized.getOrDefault(devowelized(word), "")
            }
        }
    }

    companion object {
        private val regex = Regex("[aeiou]")
        private fun lowerCase(s: String) = s.lowercase()
        private fun devowelized(s: String) = regex.replace(s.lowercase(), "#")
    }
}