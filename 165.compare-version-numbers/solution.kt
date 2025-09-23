class Solution {
    fun compareVersion(version1: String, version2: String): Int {
        val v1 = Version(version1)
        val v2 = Version(version2)
        repeat(maxOf(v1.getLength(), v2.getLength())) { i ->
            if (v1.getRevision(i) < v2.getRevision(i)) return -1
            if (v1.getRevision(i) > v2.getRevision(i)) return 1
        }
        return 0
    }
}

class Version(version: String) {
    val revisions = version.split('.').map { it.toInt() }

    fun getRevision(n: Int) = revisions.getOrElse(n) { 0 }
    fun getLength() = revisions.size
}