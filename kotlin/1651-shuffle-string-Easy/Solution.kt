class Solution {
    fun restoreString(s: String, indices: IntArray): String {
        val res = CharArray(s.length)
        for (i in s.indices)
            res.set(indices[i], s.get(i))
        return String(res)
    }
}