class Solution {
    fun bestTower(towers: Array<IntArray>, center: IntArray, radius: Int): IntArray =
            towers
                .filter { manhattan(it, center) <= radius }
                .minWithOrNull(
                    compareBy(
                        { -it[2] },
                        { it[0] },
                        { it[1] }
                    )
                )
                ?.let { intArrayOf(it[0], it[1]) }
                ?: intArrayOf(-1, -1)

    fun manhattan(a: IntArray, b: IntArray): Int = abs(a[0] - b[0]) + abs(a[1] - b[1])
}