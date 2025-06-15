class Solution {
    fun canConstruct(s: String, k: Int): Boolean {
        return s.length >= k &&
            s.groupingBy { it }
                .eachCount()
                .count { (_, v) -> v % 2 == 1 } <= k
    }
}