class Solution {
    fun romanToInt(s: String): Int {
        val symbols = mapOf('I' to 1, 'V' to 5, 'X' to 10, 'L' to 50, 'C' to 100, 'D' to 500, 'M' to 1000)
        val sequences = mapOf("IV" to 4, "IX" to 9, "XL" to 40, "XC" to 90, "CD" to 400, "CM" to 900)
        var number = 0
        var rest = s
        loop@ while (rest.isNotEmpty()) {
            for ((seq, value) in sequences) {
                if (rest.startsWith(seq)) {
                    number += value
                    rest = rest.substring(seq.length, rest.length)
                    continue@loop
                }
            }
            number += symbols.getOrDefault(rest.first(), 0)
            rest = rest.substring(1, rest.length)
        }
        return number
    }
}