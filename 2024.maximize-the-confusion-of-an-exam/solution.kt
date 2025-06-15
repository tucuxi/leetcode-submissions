class Solution {
    fun maxConsecutiveAnswers(answerKey: String, k: Int): Int {
        return listOf('T', 'F').maxOf { solve(answerKey, k, it) }
    }

    private fun solve(answerKey: String, k: Int, consecutiveAnswer: Char): Int {
        val (i, _) = answerKey.fold(0 to k) { (i, r), answer ->
            var newI = i
            var newR = r
            
            if (answer != consecutiveAnswer) {
                newR--
            }
            if (newR < 0) {
                if (answerKey[i] != consecutiveAnswer) {
                    newR++
                }
                newI++
            }

            newI to newR
        }
        return answerKey.length - i
    }
}