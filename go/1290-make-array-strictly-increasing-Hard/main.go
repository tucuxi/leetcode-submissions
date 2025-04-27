func makeArrayIncreasing(arr1 []int, arr2 []int) int {
    memo := make(map[[2]int]int)
    
    var dfs func(int, int) int
    dfs = func(i, prev int) int {
        if i == len(arr1) {
            return 0
        }
        key := [...]int{i, prev}
        if m, exists := memo[key]; exists {
            return m
        }
        cost := math.MaxInt
        if arr1[i] > prev {
            cost = dfs(i + 1, arr1[i])
        }
        k := sort.Search(len(arr2), func(i int) bool {
            return arr2[i] > prev
        })
        if k < len(arr2) {
            if cost2 := dfs(i + 1, arr2[k]); cost2 < cost {
                cost = cost2 + 1
            }
        }
        memo[key] = cost
        return cost
    }

    sort.Ints(arr2)
    unique(&arr2)
    if res := dfs(0, -1); res < math.MaxInt {
        return res
    }
    return -1
}

func unique(arr *[]int) {
    k := 1
    for i := 1; i < len(*arr); i++ {
        if (*arr)[i] > (*arr)[k - 1] {
            (*arr)[k] = (*arr)[i]
            k++
        } 
    }
    *arr = (*arr)[:k]    
}