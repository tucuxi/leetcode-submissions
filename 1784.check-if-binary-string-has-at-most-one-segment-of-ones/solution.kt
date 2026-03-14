class Solution {
    fun checkOnesSegment(s: String): Boolean {
        return s
            .dropWhile { it == '1' }
            .dropWhile { it == '0' }
            .isEmpty() 
    }
}