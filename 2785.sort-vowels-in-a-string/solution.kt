class Solution {
    fun sortVowels(s: String): String {
        val vowels = s.asIterable().filter { it.isVowel() }.sorted()
        var vowelsIterator = vowels.iterator()
        return s.map { ch ->
            if (ch.isVowel()) {
                vowelsIterator.next()
            } else {
                ch
            }
        }.joinToString("")
    }
}

fun Char.isVowel(): Boolean = "AEIOUaeiou".indexOf(this) != -1