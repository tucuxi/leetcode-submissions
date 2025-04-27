class Solution {
    fun calculate(s: String): Int {
        val stack = ArrayDeque<Int>()
        var sign = 1
        var num = 0
        var res = 0
       
        stack.push(sign)
        for (ch in s) {
	        when (ch) {
                in '0'..'9' -> {
                    num = 10 * num + (ch - '0')
                }
                '+' -> {
                    res += sign * num
                    sign = stack.peek()
                    num = 0
                }
                '-' -> {
                    res += sign * num
                    sign = -stack.peek()
                    num = 0
                }
                '(' -> {
                    stack.push(sign)
                }
                ')' -> {
                    stack.pop()
                }
            }
        }
        return res + sign * num
	}
}