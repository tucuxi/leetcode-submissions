class Solution {
    fun mctFromLeafValues(arr: IntArray): Int {
        var res = 0
        val stack = ArrayDeque<Int>().apply { add(Int.MAX_VALUE) }
        arr.forEach { a ->
            while (stack.last() <= a) {
                val mid = stack.removeLast()
                res += mid * minOf(stack.last(), a)
            }
            stack.addLast(a)
        }
        while (stack.size > 2) {
            val mid = stack.removeLast()
            res += mid * stack.last()
        }
        return res
    }
}