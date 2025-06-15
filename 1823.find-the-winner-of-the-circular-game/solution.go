func findTheWinner(n int, k int) int {
    friends := make([]int, n)
    for i := range friends {
        friends[i] = i + 1
    }
    for index := 0; len(friends) > 1; {
        index = (index + k - 1) % len(friends)
        friends = append(friends[:index], friends[index + 1:]...)
    }
    return friends[0]
}