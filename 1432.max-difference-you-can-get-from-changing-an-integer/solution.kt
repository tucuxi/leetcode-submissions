class Solution {
    fun maxDiff(num: Int): Int {
        val s = num.toString()
    
        val max = s.indexOfFirst { digit -> digit != '9' }.let { i ->
            if (i == -1)
                num 
            else
                replace(s, s[i], '9').toInt()
        }

        val min = if (s.first() != '1') {
            replace(s, s.first(), '1').toInt()
        } else {
            s.indexOfFirst { digit -> digit !in listOf('0', '1') }.let { i ->
                if (i == -1) 
                    num
                else
                    replace(s, s[i], '0').toInt()
            }
        }

        return max - min
    }

    fun replace(s: String, a: Char, b: Char): String =
        s.map { c -> if (c == a) b else c }.joinToString("")
}