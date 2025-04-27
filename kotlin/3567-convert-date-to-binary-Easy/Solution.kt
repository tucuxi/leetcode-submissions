const val SEP = "-"

class Solution {
    fun convertDateToBinary(date: String): String =
        date.split(SEP)
            .map { it.toInt().toString(2) }
            .joinToString(SEP)
}