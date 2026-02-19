class Solution {
    private var tin: IntArray = IntArray(0)
    private var tout: IntArray = IntArray(0)
    private var timer = 0
    private var up: Array<IntArray> = Array(0) { IntArray(0) }
    private var adj: Array<ArrayList<Int>> = Array(0) { ArrayList() }
    private var initialPathXor: IntArray = IntArray(0)
    private var nodeVals: IntArray = IntArray(0)
    private var bit: IntArray = IntArray(0)
    private var n = 0
    private val MAX_LOG = 17 // 50,000 nodes

    fun palindromePath(n: Int, edges: Array<IntArray>, s: String, queries: Array<String>): List<Boolean> {
        this.n = n
        adj = Array(n) { ArrayList<Int>() }
        for (edge in edges) {
            adj[edge[0]].add(edge[1])
            adj[edge[1]].add(edge[0])
        }

        tin = IntArray(n)
        tout = IntArray(n)
        initialPathXor = IntArray(n)
        nodeVals = IntArray(n)
        up = Array(n) { IntArray(MAX_LOG) }
        bit = IntArray(n + 1)

        for (i in 0 until n) {
            nodeVals[i] = 1 shl (s[i] - 'a')
        }

        dfs(0, 0, 0)

        val results = mutableListOf<Boolean>()

        for (q in queries) {
            val parts = q.split(" ")

			when(parts[0]) {

            	"update" -> {
	                val u = parts[1].toInt()
	                val charStr = parts[2]
	                val c = charStr[0]
	                val oldVal = nodeVals[u]
	                val newVal = 1 shl (c - 'a')
                
	                if (oldVal != newVal) {
	                    nodeVals[u] = newVal
	                    val diff = oldVal xor newVal
	                    updateBit(tin[u] + 1, diff) 
	                    updateBit(tout[u] + 2, diff) 
	                }
				}

				"query" -> {
	                val u = parts[1].toInt()
	                val v = parts[2].toInt()
	                val lcaNode = getLca(u, v)
	                val currentPathU = initialPathXor[u] xor queryBit(tin[u] + 1)
	                val currentPathV = initialPathXor[v] xor queryBit(tin[v] + 1)
	                val pathMask = currentPathU xor currentPathV xor nodeVals[lcaNode]
	                results.add(Integer.bitCount(pathMask) <= 1)
				}
            }
        }
        
        return results
    }

    private fun dfs(u: Int, p: Int, currentXor: Int) {
        tin[u] = timer++
        val nextXor = currentXor xor nodeVals[u]
        initialPathXor[u] = nextXor
        
        up[u][0] = p
        for (i in 1 until MAX_LOG) {
            up[u][i] = up[up[u][i - 1]][i - 1]
        }

        for (v in adj[u]) {
            if (v != p) {
                dfs(v, u, nextXor)
            }
        }
        
        tout[u] = timer - 1
    }

    private fun isAncestor(u: Int, v: Int): Boolean {
        return tin[u] <= tin[v] && tout[u] >= tout[v]
    }

    private fun getLca(u: Int, v: Int): Int {
        if (isAncestor(u, v)) return u
        if (isAncestor(v, u)) return v
        
        var curr = u
        for (i in MAX_LOG - 1 downTo 0) {
            if (!isAncestor(up[curr][i], v)) {
                curr = up[curr][i]
            }
        }
        return up[curr][0]
    }

    private fun updateBit(index: Int, value: Int) {
        var i = index
        while (i <= n) {
            bit[i] = bit[i] xor value
            i += i and -i
        }
    }

    private fun queryBit(index: Int): Int {
        var i = index
        var res = 0
        while (i > 0) {
            res = res xor bit[i]
            i -= i and -i
        }
        return res
    }
}