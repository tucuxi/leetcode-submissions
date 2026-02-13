class Solution {
    fun longestBalanced(s: String): Int {
        return maxOf(
            longestOne(s),
            longestTwo(s, 'a', 'b'),
            longestTwo(s, 'a', 'c'),
            longestTwo(s, 'b', 'c'),
            longestThree(s)
        )   
    }

    fun longestOne(s: String): Int {
        var res = 0
        var i = 0
        while (i < s.length) {
            var j = i
            while (j < s.length && s[j] == s[i]) {
                j++
            }
            res = maxOf(res, j - i)
            i = j
        }
        return res
    }

    fun longestTwo(s: String, x: Char, y: Char): Int {
        var res = 0
        var i = 0
        while (i < s.length) {
            while (i < s.length && !s[i].equals(x) && !s[i].equals(y)) {
                i++
            }
            val pos = mutableMapOf(0 to i - 1)
            var diff = 0

            while (i < s.length && (s[i] == x || s[i] == y)) {
                diff += if (s[i] == x) 1 else -1
                if (pos.contains(diff)) {
                    res = maxOf(res, i - pos.getValue(diff))
                } else {
                    pos[diff] = i
                }
                i++
            }
        }
        return res
    }

    fun longestThree(s: String): Int {
        val v = mutableMapOf((0 to 0) to -1)
        var a = 0
        var b = 0
        var c = 0
        var res = 0

        for (i in s.indices) {
            when(s[i]) {
                'a' -> a++
                'b' -> b++
                'c' -> c++
            }
            val key = (a - b) to (a - c)
            if (v.contains(key)) {
                res = maxOf(res, i - v.getValue(key))
            } else {
                v[key] = i
            }
        }
        return res
    }
}