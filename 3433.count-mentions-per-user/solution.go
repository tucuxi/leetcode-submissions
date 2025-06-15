type offlineEvent struct {
    timestamp int
    id int
}

func parseIds(s string) []int {
    var ids []int

    for len(s) > 2 {
        s = s[2:]
        id := 0
        for len(s) > 0 && s[0] != ' ' {
            id = 10 * id + int(s[0] - '0')
            s = s[1:]
        }
        if len(s) > 0 {
            s = s[1:]
        }
        ids = append(ids, id)
    }

    return ids
}

func countMentions(numberOfUsers int, events [][]string) []int {
    var (
        allEvents = 0
        hereEvents []int
        offlineEvents []offlineEvent

        mentions = make([]int, numberOfUsers)
        online = make([]int, numberOfUsers)
    )

    for _, event := range events {
        t, _ := strconv.Atoi(event[1])
        if event[0] == "MESSAGE" {
            switch event[2] {
            case "ALL":
                allEvents++
            case "HERE":
                hereEvents = append(hereEvents, t)
            default:
                for _, id := range parseIds(event[2]) {
                    mentions[id]++
                }
            }
        } else {
            id, _ := strconv.Atoi(event[2])
            offlineEvents = append(offlineEvents, offlineEvent{t, id})
        }
    }

    for id := range numberOfUsers {
        mentions[id] += allEvents
    }

    slices.Sort(hereEvents)
    slices.SortFunc(offlineEvents, func(e1, e2 offlineEvent) int { return e1.timestamp - e2.timestamp })
    
    i := 0

    for _, t := range hereEvents {
        for ; i < len(offlineEvents) && offlineEvents[i].timestamp <= t; i++ {
            online[offlineEvents[i].id] = offlineEvents[i].timestamp + 60
        }
        for id := range mentions {
            if online[id] <= t {
                mentions[id]++
            }
        }
    }

    return mentions
}