class Solution {
    fun isUgly(n: Int): Boolean {
        if (n <= 0) return false
        var p = n
        for (i in listOf(2, 3, 5)) {
            while (p % i == 0) p /= i
        }
        return p == 1 
    }
}