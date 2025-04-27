import kotlin.math.min

class Solution {
    fun distributeCandies(candyType: IntArray): Int {
        return min(candyType.size / 2, candyType.toSet().size)
    }
}