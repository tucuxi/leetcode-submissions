const val ROOT = -2
const val UNVISITED = -1

class Solution {
    fun criticalConnections(n: Int, connections: List<List<Int>>): List<List<Int>> {
        var num = Array<Int>(n) { UNVISITED } 
        var low = Array<Int>(n) { UNVISITED }
        var parents = Array<Int>(n) { UNVISITED }
        var adjList = Array<MutableList<Int>>(n) { mutableListOf() }
        var counter = 0
        val res = mutableListOf<List<Int>>()
    
        fun findBridge(u: Int) {
            num[u] = counter
            low[u] = counter
            counter++
            adjList[u].forEach { v->
                if (num[v] == UNVISITED) {
                    parents[v] = u
                    findBridge(v)
                    if (low[v] > num[u]) {
                        res.add(listOf(v, u))
                    }
                    low[u] = minOf(low[u], low[v])
                } else if (parents[u] != v) {
                    low[u] = minOf(low[u], num[v])
                }
            }
        }
    
        connections.forEach { (u, v) ->
            adjList[u].add(v)
            adjList[v].add(u)
        }
        for (i in num.indices) {
            if (num[i] == UNVISITED) {
                parents[i] = ROOT
                findBridge(i)
            }
        }
        return res
    }        
}