class Solution {
    fun largestEven(s: String): String {
        return s.dropLastWhile { it == '1' }
    }
}