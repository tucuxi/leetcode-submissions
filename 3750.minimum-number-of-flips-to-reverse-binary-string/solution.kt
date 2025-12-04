class Solution {
    fun minimumFlips(n: Int): Int {
        val s = n.toString(2)
        return (s zip s.reversed()).count { it.first != it.second }
    }
}