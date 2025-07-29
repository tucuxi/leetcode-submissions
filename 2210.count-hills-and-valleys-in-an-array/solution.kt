class Solution {
    fun countHillValley(nums: IntArray): Int {
        var res = 0
        var ascending = false
        var descending = false
        
        nums.asSequence().zipWithNext().forEach { (a, b) ->
            when {
                a < b -> {
                    ascending = true
                    if (descending) {
                        res++
                        descending = false
                    }
                }
                a > b -> {
                    descending = true
                    if (ascending) {
                        res++
                        ascending = false
                    }
                }
            }
        }

        return res
    }
}