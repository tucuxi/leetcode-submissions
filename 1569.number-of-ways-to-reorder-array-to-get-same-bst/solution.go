const modulo = 1_000_000_007

func numOfWays(nums []int) int {
	triangle := pascalTriangle(len(nums) + 1)    
	ways := calcNumOfWays(triangle, nums) - 1
	return ways % modulo
}

func calcNumOfWays(pascalTriangle [][]int, nums []int) int {
	if len(nums) <= 2 {
		return 1
	}
	left, right := getLessAndGreaterThan(nums)
	waysLeft := calcNumOfWays(pascalTriangle, left) % modulo
	waysRight := calcNumOfWays(pascalTriangle, right) % modulo
	permutations := pascalTriangle[len(left) + len(right)][len(right)] % modulo
	totalPermutations := permutations % modulo
    totalPermutations = (totalPermutations * waysLeft) % modulo
	totalPermutations = (totalPermutations * waysRight) % modulo
	return totalPermutations
}

func getLessAndGreaterThan(numbers []int) ([]int, []int) {
	var left, right []int
	number := numbers[0]
    for _, n := range numbers[1:] {
		if n < number {
			left = append(left, n)
		} else {
			right = append(right, n)
		}
	}
	return left, right
}

func pascalTriangle(length int) [][]int {
	triangle := make([][]int, length)
	for i := range triangle {
		triangle[i] = make([]int, i + 1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = (triangle[i-1][j-1] + triangle[i-1][j]) % modulo
		}
	}
	return triangle
}