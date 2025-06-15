const val first = 0
const val last = 4

class Solution {
    fun countVowelStrings(n: Int): Int {
        var res = 0
        
        fun rec(start: Int, k: Int) {
            if (k == n) {
                res++
                return
            }
            for (i in start..last) {
                rec(i, k + 1)
            }
        }
        
        rec(first, 0)
        return res
    }
}