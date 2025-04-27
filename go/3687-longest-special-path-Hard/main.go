type edge struct {
    head *node
    length int
}

type node struct {
    value int
    down []edge
}

func longestSpecialPath(edges [][]int, nums []int) []int {
    nodes := make([]*node, len(nums))
    for i := range nums {
        nodes[i] = &node{value: nums[i]}
    }

    for _, e := range edges {
        tail, head, length := nodes[e[0]], nodes[e[1]], e[2]
        head.down = append(head.down, edge{tail, length})
        tail.down = append(tail.down, edge{head, length})
    }

    prefix := make(map[*node]int)

    var prefixDfs func(*node, int)

    prefixDfs = func(u *node, acc int) {
        var actualDownEdges []edge

        prefix[u] = acc
        for _, e := range u.down {
            v := e.head
            if _, exists := prefix[v]; !exists {
                prefixDfs(v, acc + e.length)
                actualDownEdges = append(actualDownEdges, e)
            }
        }
        u.down = actualDownEdges
    }

    prefixDfs(nodes[0], 0)

    bestSolution := []int{0, math.MaxInt} 
    occ := make(map[int]int)

    var dfsLsp func(*node, []*node, int)

    dfsLsp = func(u *node, path []*node, startIndex int) {
        previousOcc := -1

        if p, exists := occ[u.value]; exists {
            previousOcc = p
        }

        startIndex = max(previousOcc + 1, startIndex)
        currentIndex := len(path)
        occ[u.value] = currentIndex
        path = append(path, u)

        if startIndex <= currentIndex {
            solution := []int{prefix[u] - prefix[path[startIndex]], currentIndex - startIndex + 1}
            if better(solution, bestSolution) {
                bestSolution = solution
            }
        }

        for _, e := range u.down {
            dfsLsp(e.head, path, startIndex)
        }

        occ[u.value] = previousOcc
    }

    dfsLsp(nodes[0], nil, 0)

    return bestSolution
}

func better(x, y []int) bool {
    return x[0] > y[0] || x[0] == y[0] && x[1] < y[1]
}