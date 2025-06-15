class Solution {
    fun numUniqueEmails(emails: Array<String>): Int =
        emails.map { email -> normalize(email) }.toSet().size

    fun normalize(email: String): String {
        val res = StringBuilder(email.length)
        var local = true
        var ignore = false
        for (ch in email) {
            if (ch == '@') {
                local = false
                ignore = false
            } else if (ch == '.' && local) {
                continue
            } else if (ch == '+' && local) {
                ignore = true
            }
            if (!ignore) {
                res.append(ch)
            }
        }
        return res.toString()
    }
}
