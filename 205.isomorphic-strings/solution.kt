class Solution {
    fun isIsomorphic(s: String, t: String): Boolean {
        if (s.length != t.length) return false
        val mapst = mutableMapOf<Char, Char>()
        val mapts = mutableMapOf<Char, Char>()
        return s.zip(t).all {
            if (mapts.contains(it.second))
                it.first == mapts.get(it.second)
            else {
                mapts.put(it.second, it.first)
                mapst.put(it.first, it.second) == null
            }
        }
    }
}