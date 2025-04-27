class Solution {
    fun shortestToChar(S: String, C: Char): IntArray {
        var occ = (0 until S.length).filter { S[it] == C }
        val res = IntArray(S.length) { if (it <= occ[0]) occ[0] - it else it - occ[0] }
        for (pos in occ.drop(1)) {
            for (i in pos downTo 0) {
                if (res[i] <= pos - i) break
                res[i] = pos - i
            }
            for (i in pos + 1 until S.length) {
                res[i] = i - pos
            }
        }
        return res
    }
}