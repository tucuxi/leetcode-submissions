class Solution {
    fun shiftingLetters(s: String, shifts: Array<IntArray>): String {
        val changes = shifts
            .flatMap { (start, end, direction) ->
                listOf(start to 2 * direction - 1, end + 1 to 1 - 2 * direction)
            }
            .groupBy { (index, _) -> index }
            .mapValues { (_, list) -> list.fold(0) { acc, (_, change) -> acc + change } }
            .toSortedMap()  

        return buildString {
            var stringIndex = 0
            var offset = 0

            changes.forEach { (index, change) ->
                s.substring(stringIndex until index).forEach { ch ->
                    append('a' + ((((ch - 'a' + offset) % 26) + 26) % 26))
                }
                stringIndex = index
                offset += change
            }
            
            append(s.substring(stringIndex))
        }
    }
}