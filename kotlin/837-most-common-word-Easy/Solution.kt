class Solution {
    fun mostCommonWord(paragraph: String, banned: Array<String>): String {
        val countMap = mutableMapOf<String, Int>()
        val bannedWords = banned.toSet()
        paragraph.split(" ", "!", "?", "'", ",", ";", ".").forEach { word ->
            val lowerCaseWord = word.toLowerCase()
            if (lowerCaseWord.isNotBlank() && !bannedWords.contains(lowerCaseWord)) {
                countMap[lowerCaseWord] = countMap[lowerCaseWord]?.plus(1) ?: 1
            }
        }
        var maxCount = 0
        var mostFrequentWord = ""
        countMap.forEach { word, count ->
            if (count > maxCount) {
                maxCount = count
                mostFrequentWord = word
            }
        }
        return mostFrequentWord
    }
}