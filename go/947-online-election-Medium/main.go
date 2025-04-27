type TopVotedCandidate struct {
    times  []int
    leader []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
    votes := make(map[int]int)
    topCandidate := 0
    maxVotes := 0
    leader := make([]int, len(persons))
    for i, p := range persons {
        votes[p]++
        if votes[p] >= maxVotes {
            maxVotes = votes[p]
            topCandidate = p
        }
        leader[i] = topCandidate
    }
    return TopVotedCandidate{times, leader}
}

func (this *TopVotedCandidate) Q(t int) int {
    k := sort.Search(len(this.times), func(i int) bool {
        return this.times[i] > t
    })
    return this.leader[k-1]
}