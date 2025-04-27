class Solution {
    fun detectCapitalUse(word: String): Boolean =
        word.all { it.isUpperCase() } || word.all { it.isLowerCase() } ||
                word.isNotEmpty() && word[0].isUpperCase() && word.substring(1).all { it.isLowerCase() }
}