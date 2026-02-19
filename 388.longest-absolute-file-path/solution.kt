class Solution {
    fun lengthLongestPath(input: String): Int {
        var maxLength = 0
        val lengths = IntArray(10000)

        for (line in input.split('\n')) {
            val level = line.lastIndexOf("\t") + 1
            val name = line.substring(level)

            if (name.contains('.')) {
                maxLength = maxOf(maxLength, lengths[level] + level + name.length)
            } else {
                lengths[level + 1] = lengths[level] + name.length
            }
        }
        return maxLength
    }
}