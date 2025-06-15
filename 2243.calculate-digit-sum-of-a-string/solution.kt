class Solution {
    tailrec fun digitSum(s: String, k: Int): String {
        return if (s.length <= k) {
            s
        } else {
            val t = s.chunked(k) {
                group -> group.map {
                    digit -> digit - '0'
                }
                .sum()
                .toString()
            }
            .joinToString(separator = "")
            digitSum(t, k)
        }
    }
}