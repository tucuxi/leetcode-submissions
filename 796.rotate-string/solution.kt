class Solution {
    fun rotateString(s: String, goal: String): Boolean {
        val n = goal.length
        if (s.length != n) return false
        outer@ for (i in 0 until n) {
            for (j in 0 until n) {
                if (s[j] != goal[(i + j) % n]) continue@outer
            }
            return true
        }
        return false
    }
}