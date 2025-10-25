class Solution {
    fun findLexSmallestString(s: String, a: Int, b: Int): String {
        val queue = mutableListOf(s)
        val visited = mutableSetOf(s)
        var res = s

        while (queue.isNotEmpty()) {
            val t = queue.removeFirst()
            res = minOf(t, res)
            accumulate(t, a).let { if (visited.add(it)) queue.add(it) }
            rotate(t, b).let { if (visited.add(it)) queue.add(it) }
        }

        return res
    }
}

fun accumulate(s: String, a: Int) =
    String(CharArray(s.length) { i -> if (even(i)) s[i] else ('0' + ((s[i] - '0') + a) % 10).toChar() })

fun rotate(s: String, b: Int) =
    String(CharArray(s.length) { i -> s[(i + b) % s.length] })

fun even(a: Int) = a % 2 == 0