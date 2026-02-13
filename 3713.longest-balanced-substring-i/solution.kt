class Solution {
    fun longestBalanced(s: String): Int {
        var res = 0

        for (i in 0 until s.length) {
            val count = IntArray(26)
            var maxCount = 0

            loop@
            for (j in i until s.length) {
                val ch = s[j] - 'a'
                count[ch]++
                if (count[ch] < maxCount) {
                    continue
                }
                maxCount = count[ch]
                for (c in count) {
                    if (c != 0 && c != maxCount) {
                        continue@loop
                    }
                }
                res = maxOf(res, j - i + 1)
            }
        }

        return res
    }
}