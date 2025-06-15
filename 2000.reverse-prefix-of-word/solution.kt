class Solution {
    fun reversePrefix(word: String, ch: Char): String {
        val pos = word.indexOf(ch)
        return if (pos == -1) {
            word
        } else {
            word.subSequence(0..pos).toList().reversed().joinToString("") + word.subSequence(pos+1 until word.length) 
        }
    }
}