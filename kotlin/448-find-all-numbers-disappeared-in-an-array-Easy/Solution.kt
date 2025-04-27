class Solution {
    fun findDisappearedNumbers(nums: IntArray): List<Int> {
        val candidates = Array<Int?>(nums.size) { it + 1 }
        nums.forEach { candidates[it - 1] = null }
        return candidates.filterNotNull()
    }
}