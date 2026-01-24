class Solution {
    fun minDeletionSize(strs: Array<String>): Int {
        val w = strs.first().length
        val dp = IntArray(w) { 1 }
        for (i in w - 2 downTo 0) {
            search@ for (j in i + 1 until w) {
                for (s in strs) {
                    if (s[i] > s[j]) {
                        continue@search
                    }
                }
                dp[i] = maxOf(dp[i], dp[j] + 1)
            }
        }

        return w - dp.max()
    }
}