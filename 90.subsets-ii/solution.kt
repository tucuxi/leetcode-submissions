class Solution {
    fun subsetsWithDup(nums: IntArray): List<List<Int>> {
        val res = mutableSetOf<List<Int>>(emptyList())
        nums.sorted().forEach { n ->
            val l = res.map { it + n }
            res.addAll(l)
        }
        return res.toList()
    }
}