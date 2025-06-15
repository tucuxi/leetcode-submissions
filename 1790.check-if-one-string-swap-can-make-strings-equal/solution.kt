class Solution {
    fun areAlmostEqual(s1: String, s2: String): Boolean {
        val d = s1.zip(s2).filter { (c1, c2) -> c1 != c2 }
        return d.isEmpty() || d.size == 2 && d[0].first == d[1].second && d[0].second == d[1].first
    }
}