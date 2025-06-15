class Solution {
    fun differenceOfSums(n: Int, m: Int): Int {
        return (1..n)
            .partition { it % m != 0 }
            .run { first.sum() - second.sum() }
    }
}