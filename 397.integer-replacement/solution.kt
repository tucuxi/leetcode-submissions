class Solution {
    fun integerReplacement(n: Int): Int {
        // Use Long to handle the edge case of Int.MAX_VALUE overflow
        var current = n.toLong()
        var count = 0
        
        while (current > 1) {
            when {
                current % 2 == 0L -> current /= 2
                current == 3L || (current % 4 == 1L) -> current--
                else -> current++
            }
            count++
        }
        
        return count
    }
}