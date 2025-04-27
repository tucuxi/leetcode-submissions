class Solution {
    fun uncommonFromSentences(s1: String, s2: String): Array<String> =
        (s1.split(" ") + s2.split(" "))
        .groupingBy { it }
        .eachCount()
        .filter { (_, n) -> n == 1 }
        .keys
        .toTypedArray()
}