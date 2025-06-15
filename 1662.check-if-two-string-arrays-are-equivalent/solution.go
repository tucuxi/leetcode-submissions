func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    i1 := newIterator(word1)
    i2 := newIterator(word2)
    for i1.hasNext() && i2.hasNext() {
        if i1.next() != i2.next() {
            return false
        }
    }
    return i1.hasNext() == i2.hasNext()
}

type iterator struct {
    word []string
    w, c int
}

func newIterator(word []string) iterator {
    return iterator{word: word}
}

func (i iterator) hasNext() bool {
    return i.w < len(i.word)
}

func (i *iterator) next() byte {
    b := i.word[i.w][i.c]
    i.c++
    if i.c == len(i.word[i.w]) {
        i.c = 0
        i.w++
    }
    return b
}