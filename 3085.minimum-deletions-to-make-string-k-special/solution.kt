class Solution {
    fun minimumDeletions(word: String, k: Int): Int {
        val h = word.groupingBy { it }.eachCount()
        
        return h.asSequence().fold(word.length) { acc, (_, a) ->
            val deleted = h.asSequence().fold(0) { acc, (_, b) ->
                if (a > b) {
                    acc + b
                } else {
                    acc + maxOf(0, b - (a + k))
                }
            }
            
            minOf(acc, deleted)
        }
    }
}