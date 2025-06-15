class Solution {
    fun canBeTypedWords(text: String, brokenLetters: String): Int {
        return text.split(' ').count { word -> 
            word.none { ch ->
                brokenLetters.contains(ch)
            }
        }
    }
}