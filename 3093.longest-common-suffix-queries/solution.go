func stringIndices(wordsContainer []string, wordsQuery []string) []int {
    t := newTrie()

    for i, w := range wordsContainer {
        t.add(w, i)
    }

    ans := make([]int, 0, len(wordsQuery))

    for _, q := range wordsQuery {
        index := t.find(q)
        ans = append(ans, index)
    }

    return ans
}

type node struct {
    index     int
    minLength int
    children  [26]*node
}

func newNode(index, length int) *node {
    return &node{index: index, minLength: length}
}

type trie struct {
    root *node
}

func newTrie() *trie {
    return &trie{root: newNode(-1, math.MaxInt)}
}

func (t *trie) add(s string, index int) {
    p := t.root
    l := len(s)

    if l < p.minLength {
        p.index = index
        p.minLength = l
    }

    for i := l-1; i >= 0; i-- {
        ch := s[i] - 'a'
        if q := p.children[ch]; q != nil {
            if l < q.minLength {
                q.index = index
                q.minLength = l
            }
            p = q
        } else {
            q = newNode(index, l)
            p.children[ch] = q
            p = q
        }
    }
}

func (t *trie) find(s string) int {
    p := t.root

    for i := len(s) - 1; i >= 0; i-- {
        ch := s[i] - 'a'
        if q := p.children[ch]; q != nil {
            p = q
        } else {
            break
        }
    }

    return p.index
}