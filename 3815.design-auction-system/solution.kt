class AuctionSystem() {

    private val userItemMap = mutableMapOf<Pair<Int, Int>, Int>()
    private val bidsByItem = mutableMapOf<Int, TreeSet<Pair<Int, Int>>>()

    fun addBid(userId: Int, itemId: Int, bidAmount: Int) {
        userItemMap
            .put(userId to itemId, bidAmount)
            ?.let { previousAmount ->
                bidsByItem.getValue(itemId).remove(userId to previousAmount)
            }
        bidsByItem
            .getOrPut(itemId) {
                TreeSet<Pair<Int, Int>>(
                    compareBy(
                        { it.second },
                        { it.first }
                    )
                )
            }
            .add(userId to bidAmount)
    }

    fun updateBid(userId: Int, itemId: Int, newAmount: Int) {
        addBid(userId, itemId, newAmount)
    }

    fun removeBid(userId: Int, itemId: Int) {
        val key = userId to itemId
        val bidAmount = userItemMap.remove(key)!!
        bidsByItem.getValue(itemId).remove(userId to bidAmount)
    }

    fun getHighestBidder(itemId: Int): Int {
        val bids = bidsByItem[itemId]
        return if (bids.isNullOrEmpty()) -1 else bids.last().first
    }

}

/**
 * Your AuctionSystem object will be instantiated and called as such:
 * var obj = AuctionSystem()
 * obj.addBid(userId,itemId,bidAmount)
 * obj.updateBid(userId,itemId,newAmount)
 * obj.removeBid(userId,itemId)
 * var param_4 = obj.getHighestBidder(itemId)
 */