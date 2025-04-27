class Solution {
    fun permuteUnique(nums: IntArray): List<List<Int>> {
        val res = mutableSetOf<List<Int>>()
        val pre = mutableListOf<Int>()
        
        fun rec(k: Int) {
            if (k == nums.size) {
                res.add(pre.map { nums[it] })
                return
            }
            for (i in nums.indices) {
                if (i !in pre) {
                    pre.add(i)
                    rec(k + 1)
                    pre.remove(i)
                }
            }
        }
        
        rec(0)
        return res.toList()
    }
}