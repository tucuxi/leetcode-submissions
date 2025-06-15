func countOrders(n int) int {
    const mod = 1_000_000_007
    res := 1
	for i := 2; i <= n; i++ {
		res = res * (2 * i - 1) * i % mod
	}
	return res
}