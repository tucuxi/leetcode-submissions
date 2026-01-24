const MOD = 1_000_000_007

func countPermutations(complexity []int) int {
    f := 1
    for i := 1; i < len(complexity); i++ {
        if complexity[i] <= complexity[0] {
            return 0
        }
        f = (f * i) % MOD
    }
    return f
}