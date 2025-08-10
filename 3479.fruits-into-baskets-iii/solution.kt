class Solution {
    fun numOfUnplacedFruits(fruits: IntArray, baskets: IntArray): Int {
        if (baskets.size == 0) {
            return fruits.size
        }
        var count = 0
        val segmentTree = SegmentTree(baskets)
        fruits.forEach { f->
            var l = 0
            var r = baskets.lastIndex
            var res = -1
            while (l <= r) {
                val mid = (l + r) / 2
                if (segmentTree.query(0, mid) >= f) {
                    res = mid
                    r = mid - 1
                } else {
                    l = mid + 1
                }
            }
            if (res != -1 && baskets[res] >= f) {
                segmentTree.update(res, Int.MIN_VALUE)
            } else {
                count++
            }
        }
        return count
    }
}

class SegmentTree(nums: IntArray) {
    val nodes = IntArray(4 * nums.size) { Int.MIN_VALUE }
    val m = nums.lastIndex

    init {
        build(nums, 1, 0, nums.lastIndex)
    }

    fun query(ql: Int, qr: Int): Int = query(1, 0, m, ql, qr)

    fun update(pos: Int, newValue: Int) = update(1, 0, m, pos, newValue)

    private fun build(nums: IntArray, p: Int, l: Int, r: Int): Int {
        nodes[p] = if (l == r) {
            nums[l]
        } else {
            val mid = (l + r) / 2
            maxOf(build(nums, 2 * p, l, mid), build(nums, 2 * p + 1, mid + 1, r))
        }
        return nodes[p]
    }

    private fun query(p: Int, l: Int, r: Int, ql: Int, qr: Int): Int {
        return when {
            ql > r || qr < l -> {
                Int.MIN_VALUE
            }
            ql <= l && r <= qr -> {
                nodes[p]
            }
            else -> { 
                val mid = (l + r) / 2
                maxOf(query(2 * p, l, mid, ql, qr), query(2 * p + 1, mid + 1, r, ql, qr))
            }
        }
    }

    private fun update(p: Int, l: Int, r: Int, pos: Int, newValue: Int) {
        if (l == r) {
            nodes[p] = newValue
        } else {
            val mid = (l + r) / 2
            if (pos <= mid) {
                update(2 * p, l, mid, pos, newValue)
            } else {
                update(2 * p + 1, mid + 1, r, pos, newValue)
            }
            nodes[p] = maxOf(nodes[2 * p], nodes[2 * p + 1])
        }
    }
}