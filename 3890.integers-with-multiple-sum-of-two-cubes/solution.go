func findGoodIntegers(n int) []int {
    var cubes []int

    for i := 1; i*i*i < n; i++ {
        cubes = append(cubes, i*i*i)
    }

    h := make(map[int]int)

    for i := range cubes {
        for j := i; j < len(cubes) && cubes[i] + cubes[j] <= n; j++ {
            h[cubes[i] + cubes[j]]++
        }
    }

    var res []int

    for num, count := range h {
        if count > 1 {
            res = append(res, num)
        }
    }

    slices.Sort(res)

    return res
}