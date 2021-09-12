func canCompleteCircuit(gas []int, cost []int) int {
	start := len(gas) - 1
	end := 0
	sum := gas[start] - cost[start]
	for start > end {
		if sum >= 0 {
			sum += gas[end] - cost[end]
			end++
		} else {
			start--
			sum += gas[start] - cost[start]
		}
	}

	if sum >= 0 {
		return start
	}

	return -1
}

// O(n^2)
func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
    for start := 0; start < n; start++ {
        // 当前储存的油气，如果到不了下一站，直接从下一站开始检测起点
        store := gas[start] - cost[start]
        if store < 0 {
            continue
        }
        station := (start+1)%n
        for station != start {
            // 计算到station下一站的油气余量
            store = store + gas[station] - cost[station]
            if store < 0 {
                break
            }
            station = (station+1)%n
        }
        // 能到达最开始的站点
        if station == start {
            return start
        }
    }
    return -1
}
