class Solution {
    fun maximumWealth(accounts: Array<IntArray>): Int {
        val wealth = accounts.map { it.sum() }
        return wealth.fold(0) { max, w -> if (w > max) w else max }
    }
}