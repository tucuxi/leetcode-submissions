class Solution {
    fun pickGifts(gifts: IntArray, k: Int): Long {
        val pq = PriorityQueue<Int>(compareByDescending { it })
        var remainingGifts = 0L
        gifts.forEach {
            pq.offer(it)
            remainingGifts += it
        }
        repeat(k) {
            val giftsFromTopPile = pq.poll()
            val giftsToPutBack = sqrt(giftsFromTopPile.toDouble()).toInt()
            pq.offer(giftsToPutBack)
            remainingGifts -= giftsFromTopPile - giftsToPutBack
        }
        return remainingGifts
    }
}