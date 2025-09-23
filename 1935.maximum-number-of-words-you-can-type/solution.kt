class Solution {
    fun canBeTypedWords(text: String, brokenLetters: String): Int {
        return text
            .split(" ")
            .count { s -> s.none { c -> c in brokenLetters } }
    }
}