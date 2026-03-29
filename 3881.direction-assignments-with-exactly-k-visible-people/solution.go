const MOD = 1_000_000_007

func countVisiblePeople(n int, pos int, k int) int {
	if k > n-1 || k < 0 {
		return 0
	}

	N := n - 1
	if k > N/2 {
		k = N - k
	}

	num := 1
	den := 1
	for i := 1; i <= k; i++ {
		num = (num * (N - i + 1)) % MOD
		den = (den * i) % MOD
	}

	ans := (num * modInverse(den)) % MOD
	return (ans * 2) % MOD
}

func power(base, exp int) int {
	res := 1
	base %= MOD
	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % MOD
		}
		base = (base * base) % MOD
		exp /= 2
	}
	return res
}

func modInverse(n int) int {
	return power(n, MOD-2)
}