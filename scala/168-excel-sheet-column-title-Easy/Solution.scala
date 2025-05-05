object Solution {
    private val K = 26
    
    def convertToTitle(n: Int): String = {
        val sb = new StringBuilder()
        var i = n
        while (i > 0) {
            i -= 1
            sb += convertToLetter(i % K)
            i /= K
        }
        sb.reverse.result
        
    }
    private def convertToLetter(m: Int): Char =
        ('A' + m).toChar
}