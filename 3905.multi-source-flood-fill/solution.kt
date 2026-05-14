class Solution {
    fun colorGrid(n: Int, m: Int, sources: Array<IntArray>): Array<IntArray> {
        val g = Array(n) { IntArray(m) }
        sources.forEach { (r, c, color) -> g[r][c] = color }

        val updates = mutableMapOf<Pair<Int, Int>, Int>()

        fun update(r: Int, c: Int, color: Int) {
            if (g[r][c] == 0) {
                val k = r to c
                updates[k] = maxOf(updates.getOrDefault(k, 0), color)
            }
        }

        val q = sources.toMutableList()

        while (q.isNotEmpty()) {
            q.forEach { (r, c, color) ->
                if (r > 0) {
                    update(r - 1, c, color)
                }
                if (r + 1 < n) {
                    update(r + 1, c, color)
                }
                if (c > 0) {
                    update(r, c - 1, color)
                }
                if (c + 1 < m) {
                    update(r, c + 1, color)
                }
            }

            q.clear()

            updates.entries.forEach { (k, color) ->
                val (r, c) = k
                g[r][c] = color
                q += intArrayOf(r, c, color)
            }

            updates.clear()
        }

        return g
    }
}