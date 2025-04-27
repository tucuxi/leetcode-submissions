class Solution {
    fun numberOfMatches(n: Int) =
        rec(0, n)
    
    private fun rec(ac: Int, n: Int): Int = 
        if (n < 2)
            ac
        else if (n % 2 == 0)
            rec(ac + n / 2, n / 2)
        else
            rec(ac + (n - 1) / 2, (n - 1) / 2 + 1)
}