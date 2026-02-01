class Solution {
    fun reverseByType(s: String): String {
        val chars = s.toCharArray()

        fun reverseByCondition(predicate: (Char) -> Boolean) {
            var left = 0
            var right = chars.lastIndex

            while (left < right) {
                while (left < right && !predicate(chars[left])) {
                    left++
                }
                while (left < right && !predicate(chars[right])) {
                    right--
                }
                if (left < right) {
                    val temp = chars[left]
                    chars[left] = chars[right]
                    chars[right] = temp
                    left++
                    right--
                }
            }
        }

        reverseByCondition { it in 'a'..'z' }
        reverseByCondition { it !in 'a'..'z' }

        return String(chars)
    }
}
