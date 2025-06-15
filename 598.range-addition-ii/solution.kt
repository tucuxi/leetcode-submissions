class Solution {
    fun maxCount(m: Int, n: Int, ops: Array<IntArray>): Int {
        val (a, b) = ops.fold(m to n) { (p, q), (a, b) -> minOf(p, a) to minOf(q, b) }
        return a * b
    }
}