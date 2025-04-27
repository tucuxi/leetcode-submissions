class Solution {
    fun vowelStrings(words: Array<String>, queries: Array<IntArray>): IntArray {
        val prefixSum = words.runningFold(0) { acc, word -> 
            if (word.isNotEmpty() && word.first() in "aeiou" && word.last() in "aeiou") {
                acc + 1
            } else {
                acc
            }
        }
        val results = queries.map { (l, r) -> prefixSum[r + 1] - prefixSum[l] }
        return results.toIntArray()
    }
}