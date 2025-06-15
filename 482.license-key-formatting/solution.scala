object Solution {
    def licenseKeyFormatting(S: String, K: Int): String = {
        val len = S.count(_.isLetterOrDigit)
        var dash = len % K
        if (dash == 0) dash = K
        val sb = new StringBuilder(len)
        for (ch <- S if ch.isLetterOrDigit) {
            if (dash == 0) {
                dash = K 
                sb += '-'
            }
            dash -= 1
            sb += ch.toUpper
        }
        sb.result
    }
}