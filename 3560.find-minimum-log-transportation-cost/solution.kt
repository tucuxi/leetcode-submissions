class Solution {
    fun minCuttingCost(n: Int, m: Int, k: Int): Long {
        val a = maxOf(0, maxOf(n, m) - k).toLong()
        return a * k.toLong() 
    }
}