class Solution {
    fun findCommonResponse(responses: List<List<String>>): String =
        responses
            .flatMap { it.distinct() }
            .groupingBy { it }
            .eachCount()
            .maxWith(compareBy<Map.Entry<String, Int>> { it.value }.thenByDescending { it.key })
            .key
}