class Solution {
    fun countPairs(words: Array<String>): Long {
        return words
            .map { s -> s
                .asIterable()
                .joinToString(separator = "") { c ->  "${'a' + (c - s[0] + 26) % 26}" }
            }
            .groupingBy { it }
            .eachCount()
            .values
            .sumOf { v -> v.toLong() * (v - 1) / 2 }
    }
}