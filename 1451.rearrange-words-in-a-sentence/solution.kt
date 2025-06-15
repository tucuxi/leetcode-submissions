class Solution {
    fun arrangeWords(text: String): String {
        return text
            .lowercase()
            .split(' ')
            .sortedBy { it.length }
            .joinToString(" ")
            .replaceFirstChar { it.uppercase() }
    }
}