class Solution {
    fun divideString(s: String, k: Int, fill: Char): Array<String> {
        return s
            .chunked(k)
            .toTypedArray()
            .apply {
                set(lastIndex, last() + "$fill".repeat(maxOf(0, k - last().length)))
            }
    }
}