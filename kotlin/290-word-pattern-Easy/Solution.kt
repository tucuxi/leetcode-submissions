class Solution {
    fun wordPattern(pattern: String, s: String): Boolean {
        val letterDict = mutableMapOf<Char, String>()
        val wordDict = mutableMapOf<String, Char>()
        val words = s.split(' ')
        if (pattern.length != words.size) {
            return false
        }
        for (i in words.indices) {
            val c = pattern[i]
            val w = words[i]
            letterDict.putIfAbsent(c, w)
            wordDict.putIfAbsent(w, c)
            if (letterDict[c] != w || wordDict[w] != c) {
                return false
            }
        }
        return true
    }
}