class Solution {
    fun numTimesAllBlue(flips: IntArray): Int {
        var q = 0
        return flips.withIndex().count { (p, flip) ->
            q = maxOf(q, flip - 1)
            q == p
        }  
    }
}