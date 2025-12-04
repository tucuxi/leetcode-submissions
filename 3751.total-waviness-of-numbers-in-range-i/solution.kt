class Solution {
    fun totalWaviness(num1: Int, num2: Int): Int =
        (maxOf(100, num1) .. num2).sumOf { it.waviness() }
}

fun Int.waviness(): Int =
    toString()
    .windowed(3)
    .count { s -> s[0] < s[1] && s[2] < s[1] || s[0] > s[1] && s[2] > s[1] }