class Solution {
    private class MaxFenwickTree(size: Int) {
        private val tree = IntArray(size + 1)

        fun updateMax(idx: Int, value: Int) {
            var i = idx
            while (i < tree.size) {
                tree[i] = maxOf(tree[i], value)
                i += i and -i
            }
        }

        fun queryMax(idx: Int): Int {
            var i = idx
            var maxVal = 0
            while (i > 0) {
                maxVal = maxOf(maxVal, tree[i])
                i -= i and -i
            }
            return maxVal
        }
    }

    fun getResults(queries: Array<IntArray>): List<Boolean> {
        val sentinel = 1 + queries.maxOf { it[1] }
        val bit = MaxFenwickTree(sentinel + 1)
        
        val obstacles = TreeSet<Int>().apply {
            add(0)
            add(sentinel)
        }

        queries.filter { it[0] == 1 }.forEach { obstacles += it[1] }

        var prevObstacle = 0

        obstacles.filter { it != 0 }.forEach {
            bit.updateMax(it, it - prevObstacle)
            prevObstacle = it
        }

        val results = mutableListOf<Boolean>()

        queries.indices.reversed().forEach { i ->
            val query = queries[i]
            val type = query[0]
            val x = query[1]

            if (type == 1) {
                val nextObstacle = obstacles.higher(x)!!
                val prevObstacle = obstacles.lower(x)!!
                bit.updateMax(nextObstacle, nextObstacle - prevObstacle)
                obstacles -= x
            } else {
                val sz = query[2]
                val lastObstacleBeforeX = obstacles.floor(x)!!
                val maxGapToLeft = bit.queryMax(lastObstacleBeforeX)
                val residualGap = x - lastObstacleBeforeX                
                results += (maxGapToLeft >= sz || residualGap >= sz)
            }
        }

        return results.reversed()
    }
}