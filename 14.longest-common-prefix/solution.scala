object Solution {
    def longestCommonPrefix(strs: Array[String]): String = {
        if (strs.isEmpty) return ""
        strs.reduce(lcp)
    }
    
    def lcp(s1: String, s2: String): String = {
        if (s2.isEmpty)
             ""
        else if (s1.startsWith(s2))
            s2
        else
            lcp(s1, s2.substring(0, s2.length - 1))
    }
}