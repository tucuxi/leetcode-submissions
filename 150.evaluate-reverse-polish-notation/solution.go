func evalRPN(tokens []string) int {
    st := new(stack)
    for _, t := range tokens {
        var k int
        switch t {
            case "+":
                k = st.pop() + st.pop()
            case "-":
                k = -st.pop() + st.pop()
            case "*":
                k = st.pop() * st.pop()
            case "/":
                a := st.pop()
                k = st.pop() / a
            default:
                k, _ = strconv.Atoi(t)
        }
        st.push(k)
    }
    return st.pop()
}

type stack struct {
    items []int
}

func (this *stack) push(a int) {
    this.items = append(this.items, a)
}

func (this *stack) pop() int {
    res := this.items[len(this.items) - 1]
    this.items = this.items[:len(this.items) - 1]
    return res
}