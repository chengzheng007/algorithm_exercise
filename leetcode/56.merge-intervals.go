import "sort"

func merge(intervals [][]int) [][]int {
	var res [][]int
	// 没有明确说不让排序的，可以先排序
	sort.Sort(TwoDimArr(intervals))

	for _, pair := range intervals {
		// res为空，或已有集合末尾的区间的右值<新区间的左值，直接插入到res末尾
		if len(res) == 0 || res[len(res)-1][1] < pair[0] {
			res = append(res, []int{pair[0], pair[1]})
		} else {
			// 将res末尾的区间的右值改为res末尾和新区间的右值中的较大者
			if res[len(res)-1][1] < pair[1] {
				res[len(res)-1][1] = pair[1]
			}
		}
	}

	return res
}

type TwoDimArr [][]int

func (l TwoDimArr) Len() int           { return len(l) }
func (l TwoDimArr) Less(i, j int) bool { return l[i][0] < l[j][0] }
func (l TwoDimArr) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }