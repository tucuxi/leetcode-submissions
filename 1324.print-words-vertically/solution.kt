class Solution {
    fun printVertically(s: String): List<String> {
        val words = s.split(' ')
        val rows = words.map { it.length }.max()
        return (0 until rows).map { row ->
            words
                .map { it.getOrElse(row) { ' ' } }
                .joinToString("")
                .trimEnd()
        }
    }
}