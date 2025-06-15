class Solution {
    fun reverseDegree(s: String): Int {
        return s.withIndex().sumOf { (i, c) -> (26 - (c - 'a')) * (i + 1) } 
    }
}