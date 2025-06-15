class Solution {
    fun capitalizeTitle(title: String): String {
        return title.split(' ').map { capitalizeWord(it) }.joinToString(" ")
    }
    
    fun capitalizeWord(word: String): String {
        return if (word.length <= 2) {
            word.toLowerCase()
        } else {
            word.first().toUpperCase() + word.drop(1).toLowerCase()
        }
    }
}