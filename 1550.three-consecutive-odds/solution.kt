class Solution {
    fun threeConsecutiveOdds(arr: IntArray): Boolean {
        return arr
            .scan(0) { acc, a -> (acc + a % 2) * (a % 2) }
            .any { it >= 3 }
    }
}