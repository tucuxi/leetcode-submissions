class Solution {
    fun shortestCompletingWord(licensePlate: String, words: Array<String>): String {
        val plateHist = countLetters(licensePlate)
        var shortestWord: String? = null
        var shortestWordLength = Int.MAX_VALUE
        words.forEach { w->
            val wordHist = countLetters(w)
            if (('a'..'z').all { wordHist[it - 'a'] >= plateHist[it - 'a'] } ) {
                if (w.length < shortestWordLength) {
                    shortestWord = w
                    shortestWordLength = w.length
                }
            }
        }
        return shortestWord ?: ""
    }
    
    private fun countLetters(s: String): IntArray {
        val h = IntArray(26)
        s.forEach { c -> 
            val cl = c.toLowerCase()
            if (cl in 'a'..'z') { h[cl - 'a']++ }
        }
        return h
    }
}