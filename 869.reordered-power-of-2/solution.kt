class Solution {
    fun reorderedPowerOf2(n: Int): Boolean {
        val digits = n.toString().toList().map { ch -> ch - '0' }
        val visited = BooleanArray(digits.size)

        fun dfs(acc: Int, remaining: Int, leading: Boolean): Boolean {
            if (remaining == 0) {
                return  isPowerOf2(acc)
            }
            var res = false
            for (i in digits.indices) {
                if (digits[i] == 0 && leading) {
                    continue
                }
                if (!visited[i]) {
                    visited[i] = true
                    res = res or dfs(10 * acc + digits[i], remaining - 1, false)
                    visited[i] = false
                }
            }
            return res
        }

        return dfs(0, digits.size, true)
    }

    fun isPowerOf2(n: Int) = n and (n - 1) == 0
}