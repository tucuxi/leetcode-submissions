import java.util.Scanner

class Solution {
    fun compareVersion(version1: String, version2: String): Int {
        val scanner1 = Scanner(version1).useDelimiter("\\.")
        val scanner2 = Scanner(version2).useDelimiter("\\.")

        while (scanner1.hasNextInt() || scanner2.hasNextInt()) {
            val v1 = if (scanner1.hasNextInt()) scanner1.nextInt() else 0
            val v2 = if (scanner2.hasNextInt()) scanner2.nextInt() else 0
            when {
                v1 < v2 -> return -1
                v1 > v2 -> return 1
            }
        }
        return 0
    }
}
