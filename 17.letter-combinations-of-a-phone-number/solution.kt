class Solution {
    companion object {
        val digitMap =
            mapOf(
                '2' to "abc", '3' to "def", '4' to "ghi", '5' to "jkl",
                '6' to "mno", '7' to "pqrs", '8' to "tuv", '9' to "wxyz"
            )
    }

    fun letterCombinations(digits: String): List<String> {
        val res = mutableListOf<String>()
        
        fun recurse(prefix: StringBuilder, index: Int) {
            if (index == digits.length) {
                res.add(prefix.toString())
                return
            }
            for (ch in digitMap[digits[index]]!!) {
                prefix.append(ch)
                recurse(prefix, index + 1)
                prefix.setLength(index)
            }
        }
        
        if (digits.isNotEmpty()) {
            recurse(StringBuilder(5), 0)
        }
        return res
    }
}