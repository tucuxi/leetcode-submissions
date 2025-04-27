func findChampion(grid [][]int) int {
    for i, g := range grid {
        strongest := true
        for j, x := range g {
            if i != j && x == 0 {
                strongest = false 
                break
            }
        }
        if strongest {
            return i
        }
    }
    return -1
}