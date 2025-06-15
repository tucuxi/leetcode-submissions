object Solution {
    def reverseOnlyLetters(S: String): String = {
        val reversedLetters = S.filter(_.isLetter).reverseIterator
        val res = new StringBuilder
        for (s <- S)
            if (s.isLetter) res += reversedLetters.next else res += s
        res.toString
    }
}