class Solution {
    fun countPrefixSuffixPairs(words: Array<String>): Int {
        var res = 0
        for (j in words.indices) {
            for (i in 0 until j) {
                if (words[j].startsWith(words[i]) && words[j].endsWith(words[i])) {
                    res++
                }
            }
        }
        return res
    }
}