class Solution {
    fun maxProduct(words: Array<String>): Int {
        val chars = words.map { w -> w.fold(0) { a, c -> a or (1 shl (c - 'a')) }}
        var maxlen = 0
        for (i in 0 until words.lastIndex) {
            for (j in i + 1 .. words.lastIndex) {
                if (chars[i] and chars[j] == 0) {
                    maxlen = maxOf(maxlen, words[i].length * words[j].length)
                }
            }
        }
        return maxlen
    }
}