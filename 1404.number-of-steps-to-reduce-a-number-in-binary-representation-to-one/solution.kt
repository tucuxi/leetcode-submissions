class Solution {
    fun numSteps(s: String): Int {
        var steps = 0
        var carry = 0

        for (i in s.lastIndex downTo 1) {
            val digit = s[i] - '0' + carry
            if (digit == 1) {
                steps++
                carry = 1
            }
            steps++
        }

        return steps + carry
    }
}