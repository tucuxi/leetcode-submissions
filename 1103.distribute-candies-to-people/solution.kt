import kotlin.math.min

class Solution {
    fun distributeCandies(candies: Int, num_people: Int): IntArray {
        val res = IntArray(num_people)
        var remainingCandies = candies
        var dueCandies = 1
        var nextPerson = 0
        while (remainingCandies > 0) {
            val givenCandies = min(dueCandies, remainingCandies)
            res[nextPerson] += givenCandies
            remainingCandies -= givenCandies
            dueCandies += 1
            nextPerson = (nextPerson + 1) % num_people
        }
        return res
    }
}