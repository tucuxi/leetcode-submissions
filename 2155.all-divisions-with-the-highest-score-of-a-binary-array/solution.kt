class Solution {
    fun maxScoreIndices(nums: IntArray): List<Int> {
        var score = nums.count { it == 1 }
        var maxScore = score
        val res = mutableListOf(0)

        nums.forEachIndexed { i, num ->
            when (num) {
                0 -> score++
                1 -> score--
            }
            if (score == maxScore) {
                res.add(i + 1)
            } else if (score > maxScore) {
                res.clear()
                res.add(i + 1)
                maxScore = score
            }
        }

        return res
    } 
}