class Solution {
    fun reverseStr(s: String, k: Int): String =
        s.chunked(k).mapIndexed { index, part -> if (index % 2 == 0) part.reversed() else part }.joinToString("")
}