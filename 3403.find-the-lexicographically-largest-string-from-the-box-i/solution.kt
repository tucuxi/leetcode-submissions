class Solution {
    fun answerString(word: String, numFriends: Int): String {
        if (numFriends == 1) {
            return word
        }
        var res = ""
        for (i in word.indices) {
            val s = word.substring(i, minOf(word.length, i + word.length - numFriends + 1))
            if (s > res) {
                res = s
            }
        }
        return res
    }
}