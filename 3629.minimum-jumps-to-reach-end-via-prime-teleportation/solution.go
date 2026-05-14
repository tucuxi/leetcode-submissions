const maxValue = 1000000

var spf []int

func init() {
    spf = make([]int, maxValue+1)
    for i := 2; i <= maxValue; i++ {
        spf[i] = i
    }
    for i := 2; i*i <= maxValue; i++ {
        if spf[i] == i {
            for j := i * i; j <= maxValue; j += i {
                if spf[j] == j {
                    spf[j] = i
                }
            }
        }
    }
}

func getPrimeFactors(num int) []int {
    var factors []int
    for num > 1 {
        p := spf[num]
        factors = append(factors, p)
        for num%p == 0 {
            num /= p
        }
    }
    return factors
}

func minJumps(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return 0
    }

    edges := make(map[int][]int)

    for i, num := range nums {
        for _, p := range getPrimeFactors(num) {
            edges[p] = append(edges[p], i)
        }
    }

    res := 0
    visited := make([]bool, n)
    visited[0] = true

    for q := []int{0}; len(q) > 0; {
        var q2 []int
        for _, i := range q {
            if i == n-1 {
                return res
            }
            
            if j := i - 1; j >= 0 && !visited[j] {
                visited[j] = true
                q2 = append(q2, j)
            }
            
            if j := i + 1; j < n && !visited[j] {
                visited[j] = true
                q2 = append(q2, j)
            }
            
            val := nums[i]
            if val >= 2 && spf[val] == val {
                if list, exists := edges[val]; exists {
                    for _, j := range list {
                        if !visited[j] {
                            visited[j] = true
                            q2 = append(q2, j)
                        }
                    }

                    delete(edges, val) 
                }
            }
        }
        q = q2
        res++
    }
    return -1
}