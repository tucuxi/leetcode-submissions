object Solution {
    def isAnagram(s: String, t: String): Boolean = {
        (s.length == t.length) && s.sorted == t.sorted
    }
}