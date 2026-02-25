class Solution {
    fun mergeAdjacent(nums: IntArray): List<Long> {
        val res = mutableListOf<Long>()
        nums.forEach { num ->
            res.add(num.toLong())
            while (res.size > 1 && res[res.lastIndex - 1] == res.last()) {
                res.removeAt(res.lastIndex)
                res[res.lastIndex] *= 2
            }
        }
        return res.toList()
    }
}