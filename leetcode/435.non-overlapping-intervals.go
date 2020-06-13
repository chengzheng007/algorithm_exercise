import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	// 找出当前区间中，最多有几个互不相交的区间
	// [1,2]、[2,3]也算互不相交
	// 那么最小需移除的区间数，等于总数减去这个数字
	// 至少有一个区间没有相交（都移除只剩一个区间时）

	if len(intervals) <= 1 {
		return 0
	}

	// 每个区间按照【结束时间】排序
	sort.Sort(intervalEndSort(intervals))

	maxNonOverlap := 1
	end := intervals[0][1]
	for _, interval := range intervals {
		// interval[0]:当前遍历区间的start
		// 比end大说明与当前区间不相交
		if interval[0] >= end {
			maxNonOverlap++
			end = interval[1]
		}
	}
	return len(intervals) - maxNonOverlap
}

type intervalEndSort [][]int

func (l intervalEndSort) Len() int           { return len(l) }
func (l intervalEndSort) Less(i, j int) bool { return l[i][1] < l[j][1] }
func (l intervalEndSort) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
