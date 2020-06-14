import "sort"

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return n
	}

	// 类比区间调度，如果有n个区间不重合，那么最少需要n支箭
	// 即转换为求最少可看成多少个不重合区间，这个数即射箭数
	// 这里不重合，start>end边界值相等，
	sort.Sort(sortByEnd(points))

	end := points[0][1]
	cnt := 1
	for i := 0; i < n; i++ {
		if points[i][0] > end {
			cnt++
			end = points[i][1]
		}
	}

	return cnt
}

type sortByEnd [][]int

func (l sortByEnd) Len() int {
	return len(l)
}
func (l sortByEnd) Less(i, j int) bool {
	return l[i][1] < l[j][1]
}
func (l sortByEnd) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}