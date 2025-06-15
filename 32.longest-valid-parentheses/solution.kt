class Solution {
    fun longestValidParentheses(s: String): Int {
        var left = 0
        var right = 0
        var maxLength = 0
        for (i in s.indices) {
            if (s[i] == '(') {
                left++
            } else {
                right++
            }
            if (left == right) {
                maxLength = maxOf(maxLength, 2 * right)
            } else if (right >= left) {
                right = 0
                left = right
            }
        }
        left = 0
        right = 0
        for (i in s.length - 1 downTo 0) {
            if (s[i] == '(') {
                left++
            } else {
                right++
            }
            if (left == right) {
                maxLength = maxOf(maxLength, 2 * left)
            } else if (left >= right) {
                right = 0
                left = right
            }
        }
        return maxLength
    }
}