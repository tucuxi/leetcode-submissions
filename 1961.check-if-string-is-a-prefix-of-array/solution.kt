class Solution {
    fun isPrefixString(s: String, words: Array<String>): Boolean {
        var start = 0
        var len = s.length
        for (w in words) {
            if (s.length - start < w.length
                    || s.subSequence(start, start + w.length) != w) {
                return false
            }
            start += w.length
            if (s.length - start == 0) {
                return true
            }
        }
        return false
    }
}