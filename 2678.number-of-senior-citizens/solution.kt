class Solution {
    fun countSeniors(details: Array<String>) =
        details
            .map { s -> 10 * (s[11] - '0') + (s[12] - '0') }
            .count { it > 60 }
}