type sample struct {
    petalLength float64
    species string
}

func calculateMaxInfoGain(petal_length []float64, species []string) float64 {
    n := float64(len(species))
    sl := sampleList(petal_length, species)
    fl, fr := map[string]int{}, map[string]int{}
    for _, s := range sl {
        fr[s.species]++
    }
    hMin := h(n, fr)
    for i := range sl {
        name := sl[i].species
        fl[name]++
        fr[name]--
        if fr[name] == 0 {
            delete(fr, name)
        }
        size := float64(i + 1)
        hl := h(size, fl) * size
        hr := h(n - size, fr) * (n - size)
        hMin = min(hMin, (hl + hr) / n)
    }
    return h(n, fl) - hMin
}

func h(n float64, f map[string]int) float64 {
    res := 0.
    for _, c := range f {
        p := float64(c) / n
        res -= p * math.Log2(p)
    }
    return res
}

func sampleList(petalLength []float64, species []string) []sample {
    res := make([]sample, len(species))
    for i := range species {
        res[i] = sample{petalLength[i], species[i]}
    }
    sort.Slice(res, func(i, j int) bool { return res[i].petalLength < res[j].petalLength })
    return res
}

func min(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}

func max(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}