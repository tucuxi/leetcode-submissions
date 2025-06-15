object Solution {
    def canConstruct(ransomNote: String, magazine: String): Boolean = {
        def histogram(s: String): Array[Int] = {
            val h = new Array[Int]('z' - 'a' + 1)
            for (c <- s) h(c - 'a') += 1
            h
        }
        val histRansom = histogram(ransomNote)
        val histMagazine = histogram(magazine)
        for ((h1, h2) <- histRansom zip histMagazine)
            if (h1 > h2) return false
        true
    }
}