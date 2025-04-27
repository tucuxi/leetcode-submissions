type trieNode struct {
    children [26]*trieNode
    count int
}

type trie struct {
    root trieNode
}

func (t *trie) insert(s string) {
    p := &t.root
    for _, r := range s {
        ch := r - 'a'
        if p.children[ch] == nil {
            p.children[ch] = new(trieNode)
        }
        p = p.children[ch]
        p.count++
    }
}

func (t *trie) sumPrefixCounts(s string) int {
    res := 0
    p := &t.root
    for _, r := range s {
        p = p.children[r - 'a']
        res += p.count
    }
    return res
}

func sumPrefixScores(words []string) []int {
    t := new(trie)
    for _, w := range words {
        t.insert(w)
    }
    res := make([]int, 0, len(words))
    for _, w := range words {
        res = append(res, t.sumPrefixCounts(w))
    }
    return res
}