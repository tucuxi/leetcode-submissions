object Solution {
    def checkRecord(s: String): Boolean = {
        s.count(_ == 'A') < 2 && !s.contains("LLL")
    }
}