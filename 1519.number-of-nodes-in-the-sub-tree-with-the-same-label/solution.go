func countSubTrees(n int, edges [][]int, labels string) []int {
    adj := make([][]int, n)
    for _, e := range edges {
        adj[e[0]] = append(adj[e[0]], e[1])
        adj[e[1]] = append(adj[e[1]], e[0])
    }
    ans := make([]int, n)

    var traverse func(int, int) [26]int
    traverse = func(node int, parent int) [26]int {
        total := [26]int{}
        for _, child := range adj[node] {
            if child != parent {
                count := traverse(child, node)
                for i := range count {
                    total[i] += count[i]
                }
            }
        }
        total[labels[node] - 'a']++
        ans[node] = total[labels[node] - 'a']
        return total
        
    }

    traverse(0, -1)
    return ans
}