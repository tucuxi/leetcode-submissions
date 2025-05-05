class Solution {
    fun numEquivDominoPairs(dominoes: Array<IntArray>): Int =
        dominoes
            .groupingBy { (a, b) -> if (a < b) Pair(a, b) else Pair(b, a) }
            .eachCount()
            .map { (_, count) -> (count - 1) * count / 2 }
            .sum()
}