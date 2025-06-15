class Solution {
    fun kthGrammar(n: Int, k: Int): Int {
        return if (n == 1) {
            0
        } else {
            val p = kthGrammar(n - 1, (k + 1) / 2)
            if (k % 2 == 0) {
                1 - p
            } else {
                p
            }
        }
    }
}