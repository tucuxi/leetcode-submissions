class Solution {
    fun combinationSum3(k: Int, n: Int): List<List<Int>> {
        val res = mutableListOf<List<Int>>()
        val pre = mutableListOf<Int>()
        
        fun rec(start: Int, sum: Int) {
            if (sum == n && pre.size == k) {
                res.add(pre.toList())
                return
            }
            if (sum > n || pre.size >= k) {
                return
            }
            for (i in start..9) {
                 pre.add(i)
                 rec(i + 1, sum + i)
                 pre.remove(i)
            }
        }
        
        rec(1, 0)
        return res
    }
}