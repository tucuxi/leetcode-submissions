class Solution {
    fun minGenerations(points: Array<IntArray>, target: IntArray): Int {
        val gen = points.map { Point(it[0], it[1], it[2]) }.toMutableSet()
        val t = Point(target[0], target[1], target[2])
        var k = 0

        while (t !in gen) {
            val nextGen = mutableSetOf<Point>()
            for (p1 in gen) {
                for (p2 in gen) {
                    if (p1 != p2) {
                        nextGen += Point((p1.x + p2.x) / 2, (p1.y + p2.y) / 2, (p1.z + p2.z) / 2)
                    }
                }
            }
            val n = gen.size
            gen += nextGen
            if (gen.size == n) {
                return -1
            }
            k++
        }

        return k
    }

    data class Point(val x: Int, val y: Int, val z: Int)
}