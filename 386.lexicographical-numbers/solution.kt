class Solution {
    fun lexicalOrder(n: Int): List<Int> {
        val res = mutableListOf<Int>()

        fun dfs(k: Int) {
            if (k <= n) {
                res.add(k)
                repeat(10) { i -> dfs(10 * k + i) }
            }
        }

        for (i in 1..9) {
            dfs(i)
        }

        return res
    }
}