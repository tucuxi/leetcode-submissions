class Solution {
    fun minimumTeachings(n: Int, languages: Array<IntArray>, friendships: Array<IntArray>): Int {
        val nonCommunicatingFriends = friendships
            .filter { (a, b) -> languages[a - 1].toSet().intersect(languages[b - 1].toSet()).isEmpty() }
            .flatMap { (a, b) -> listOf(a - 1, b - 1) }
            .distinct()

        val speakersPerLanguage = nonCommunicatingFriends
            .flatMap { languages[it].toList() }
            .groupingBy { it }
            .eachCount()

        return nonCommunicatingFriends.size - (speakersPerLanguage.maxOfOrNull { it.value } ?: 0)
    }
}