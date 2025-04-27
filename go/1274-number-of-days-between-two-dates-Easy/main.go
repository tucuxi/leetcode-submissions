func daysBetweenDates(date1 string, date2 string) int {
    const layout = "2006-01-02"
    t1, _ := time.Parse(layout, date1)
    t2, _ := time.Parse(layout, date2)
    duration := t2.Sub(t1)
    days := math.Abs(math.Round(duration.Hours() / 24))
    return int(days)
}
