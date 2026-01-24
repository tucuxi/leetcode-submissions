class Solution {
    fun residuePrefixes(s: String): Int {
        val letters = mutableSetOf<Char>()
        var res = 0

        s.forEachIndexed { i, ch ->
            letters.add(ch)
            if (letters.size == (i + 1) % 3) {
                res++
            }
        }

        return res
    }
}