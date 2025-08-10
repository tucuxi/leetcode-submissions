class Solution {
    fun reorderedPowerOf2(n: Int): Boolean {
        val permutations = sequence<Int> {
            val digits = n.toString().map { ch -> ch - '0' }
            val visited = BooleanArray(digits.size)

            suspend fun SequenceScope<Int>.dfs(acc: Int, remaining: Int) {
                if (remaining == 0) {
                    yield(acc)
                    return
                }
                for (i in digits.indices) {
                    if (!visited[i] && (digits[i] != 0 || acc != 0)) {
                        visited[i] = true
                        dfs(10 * acc + digits[i], remaining - 1)
                        visited[i] = false
                    }
                }
            }

            dfs(0, digits.size)
        }
        return permutations.any { isPowerOf2(it) }
    }

    fun isPowerOf2(n: Int) = n and (n - 1) == 0
}