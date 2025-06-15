class Solution {
    fun findJudge(n: Int, trust: Array<IntArray>): Int {
        val trustsAnybody = BooleanArray(n + 1)
        val trustedBy = IntArray(n + 1)
        trust.forEach { (a, b) ->
            trustsAnybody[a] = true
            trustedBy[b]++
        }
        return (1..n).find { !trustsAnybody[it] && trustedBy[it] == n - 1 } ?: -1
    }
}