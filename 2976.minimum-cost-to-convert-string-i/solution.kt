const val LETTERS = 26

class Solution {
    fun minimumCost(source: String, target: String, original: CharArray, changed: CharArray, cost: IntArray): Long {
        val distance = Array(LETTERS) { IntArray(LETTERS) { Int.MAX_VALUE } }

        for (i in original.indices) {
            val x = original[i] - 'a'
            val y = changed[i] - 'a'
            distance[x][y] = minOf(distance[x][y], cost[i])
        }
        for (i in 0 until LETTERS) {
            distance[i][i] = 0
        }
        for (k in 0 until LETTERS) {
            for (i in 0 until LETTERS) {
                if (distance[i][k] < Int.MAX_VALUE) {
                    for (j in 0 until LETTERS) {
                        if (distance[k][j] < Int.MAX_VALUE) {
                            distance[i][j] = minOf(distance[i][j], distance[i][k] + distance[k][j])
                        }
                    }
                }
            }
        }

        var res = 0L

        for (i in source.indices) {
            val d = distance[source[i] - 'a'][target[i] - 'a']
            if (d == Int.MAX_VALUE) {
                return -1
            } 
            res += d.toLong()
        }

        return res
    }
}