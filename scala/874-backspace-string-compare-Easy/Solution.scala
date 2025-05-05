object Solution {
    def backspaceCompare(S: String, T: String): Boolean = {
        edit(S) == edit(T)
    }
    private def edit(s: String): String = {
        var b = new StringBuilder()
        var skip = 0
        for (i <- s.length - 1 to 0 by -1) {
            if (s(i) == '#')
                skip += 1
            else if (skip > 0)
                skip -= 1
            else
                b += s(i)     
        }
        b.reverse.result
    }
}