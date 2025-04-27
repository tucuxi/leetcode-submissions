import java.security.SecureRandom

const val RADIX = 36
const val PREFIX = "http://u.rl/"

class Codec() {
    private val map = mutableMapOf<Int, String>()
    private val gen = SecureRandom()

    // Encodes a URL to a shortened URL.
    @Synchronized fun encode(longUrl: String): String {
        while (true) {
            val rnd = gen.nextInt(1_000_000_000)
            if (rnd !in map) {
                map[rnd] = longUrl
                return PREFIX + rnd.toString(RADIX)
            }
        }        
    }

    // Decodes a shortened URL to its original URL.
    @Synchronized fun decode(shortUrl: String): String {
        val num = shortUrl.substring(PREFIX.length).toInt(RADIX)
        return map[num] ?: ""
    }
}

/**
 * Your Codec object will be instantiated and called as such:
 * var obj = Codec()
 * var url = obj.encode(longUrl)
 * var ans = obj.decode(url)
 */