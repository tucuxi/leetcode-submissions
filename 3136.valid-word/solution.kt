class Solution {
    fun isValid(word: String): Boolean {
        return word.length >= 3 &&
            word.all { c -> c.isLetterOrDigit() } &&
            word.any { c -> c.isVowel() } &&
            word.any { c -> c.isLetter() && !c.isVowel() }
    }
}

fun Char.isVowel() = "aeiouAEIOU".indexOf(this) >= 0