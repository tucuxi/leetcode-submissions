func minimumDistance(word string) int {
    memo := make(map[[3]int]int)

    var dp func(string, int, byte, byte) int

    dp = func(word string, index int, f1, f2 byte) int {
        if index == len(word) {
            return 0
        }

        key := [3]int{index, int(f1), int(f2)}
        
        if res, ok := memo[key]; ok {
            return res
        }

        res := min(
            dp(word, index+1, word[index], f2) + dist(f1, word[index]),
            dp(word, index+1, f1, word[index]) + dist(f2, word[index]),
        )

        memo[key] = res

        return res
    }

    return dp(word, 0, 0, 0)
}

func dist(a, b byte) int {
    if a == 0 {
        return 0 
    }

    xa, ya := coord(a)
    xb, yb := coord(b)

    return abs(xa - xb) + abs(ya - yb)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func coord(b byte) (int, int) {
    n := int(b - 'A')
    return n%6, n/6
}