class Solution {
    fun hasAllCodes(s: String, k: Int): Boolean {
        val codes = BooleanArray(1 shl k)
        s.windowed(k).forEach { codes[it.toInt(2)] = true }
        return codes.all { it }
    }
}