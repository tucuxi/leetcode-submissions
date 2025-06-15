object Solution {
    def isLongPressedName(name: String, typed: String): Boolean = {
        var i = 0
        var j = 0
        while (i < name.length && j < typed.length) {
            if (name(i) == typed(j)) {
                i += 1
            } else if (j == 0 || typed(j - 1) != typed(j)) {
                return false
            }
            j += 1
        }
        i == name.length
    }
}