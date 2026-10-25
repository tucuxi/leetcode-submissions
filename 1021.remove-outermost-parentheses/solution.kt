class Solution {
    fun removeOuterParentheses(s: String): String {
        return buildString {
            var k = 0
            s.forEach { ch ->
                if (ch == ')') {
                    k--
                }
                if (k > 0) {
                    append(ch)
                }
                if (ch == '(') {
                    k++
                }
            }
        }
    }
}