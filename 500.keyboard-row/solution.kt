class Solution {
    private val rows = arrayOf("qwertyuiop", "asdfghjkl", "zxcvbnm")

    fun findWords(words: Array<String>): Array<String> {
        return words.filter { word ->
            val lowerCaseWord = word.toLowerCase()
            rows.any { row ->
                lowerCaseWord.all { ch ->
                    ch in row
                }
            }
        }.toTypedArray()
    }
}