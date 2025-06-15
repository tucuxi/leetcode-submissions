class Solution {
    fun reverseOnlyLetters(s: String): String {
        val itr = s.filter { it.isLetter() }.reversed().iterator()
        val res = StringBuilder()
        s.forEach {
            val c = if (it.isLetter()) itr.next() else it
            res.append(c)
        }
        return res.toString()
    }
}