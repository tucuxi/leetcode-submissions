class Solution {
    fun maxScore(s: String): Int {
        val initialScore = s.count { it == '1' }
        return s
            .dropLast(1)
            .runningFold(initialScore) { score, ch -> score + (if (ch == '0') 1 else -1) }
            .drop(1)
            .max()
    }
}