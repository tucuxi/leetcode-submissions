class Solution {
    fun makeFancyString(s: String): String {
        return buildString {
            var c1: Char = '\u0000'
            var c2: Char = '\u0000'

            s.forEach { c ->
                if (c != c1 || c != c2) {
                    append(c)
                }
                c1 = c2
                c2 = c
            }
        }
    }
}