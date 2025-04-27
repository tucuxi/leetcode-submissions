class Solution {
    fun filterRestaurants(restaurants: Array<IntArray>, veganFriendly: Int, maxPrice: Int, maxDistance: Int): List<Int> {
        return restaurants
            .filter { (it[2] == 1 || veganFriendly == 0) && it[3] <= maxPrice && it[4] <= maxDistance }
            .sortedWith(compareBy({ -it[1] }, { -it[0] }))
            .map { it[0] }
    }
}