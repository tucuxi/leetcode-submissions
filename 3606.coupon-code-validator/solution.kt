class Solution {
    fun validateCoupons(code: Array<String>, businessLine: Array<String>, isActive: BooleanArray): List<String> {
        return code
            .indices
            .filter { i -> isActive[i] && isValidCode(code[i]) && isValidBusinessLine(businessLine[i]) }
            .sortedWith(
                compareBy(
                    { i -> businessLine[i] },
                    { i -> code[i] }
                )
            )
            .map { i -> code[i] }
    }

    fun isValidCode(code: String): Boolean =
        code.isNotEmpty() && code.all { c -> c.isLetterOrDigit() || c == '_' }

    fun isValidBusinessLine(businessLine: String): Boolean =
        businessLine in listOf("electronics", "grocery", "pharmacy", "restaurant")
}