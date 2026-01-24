class Solution {
    fun maximalRectangle(matrix: Array<CharArray>): Int {
        val n = matrix.first().size
        val height = IntArray(n)
        var res = 0

        matrix.forEach { row ->
            for (i in 0 until n) {
                height[i] = if (row[i] == '1') height[i] + 1 else 0
            }
            val stack = Stack<Int>()
            for (i in 0..n) {
                val currentHeight = if (i == n) 0 else height[i]
                while (stack.isNotEmpty() && height[stack.peek()] > currentHeight) {
                    val h = height[stack.pop()]
                    var w = if (stack.isEmpty()) i else i - stack.peek() - 1
                    res = maxOf(res, h * w)
                }
                stack.push(i)
            }
        }

        return res
    }
}