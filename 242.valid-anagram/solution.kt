class Solution {
    fun isAnagram(s: String, t: String) = s.groupBy{ it } == t.groupBy{ it }
}