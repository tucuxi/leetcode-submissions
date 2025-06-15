object Solution {
    def isPalindrome(x: Int): Boolean = {
        (x >= 0) && palin(x.toString)
    }
    def palin(s: String): Boolean = {
        if (s.length < 2) {
            true
        } else if (s.charAt(0) == s.charAt(s.length - 1)) {
            palin(s.substring(1, s.length - 1))
        } else {
            false
        }
    }
}