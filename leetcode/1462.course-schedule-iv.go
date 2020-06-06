import "container/list"

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	// [1,0]先学1后学0，构造由1指向0的有向图，0的出度为1
	// 如0->1->2->3->4,表示先学0，后学1，在学2
	outDegree := make([]int, n)
	// 构建逆邻接表reverseAdj[i]=j表示先学j后学i
	reverseAdj := make([][]int, n)

	for _, course := range prerequisites {
		outDegree[course[0]]++
		reverseAdj[course[1]] = append(reverseAdj[course[1]], course[0])
	}

	q := list.New()
	// 将没有后学课程的课放入队列
	for course, num := range outDegree {
		if num == 0 {
			q.PushBack(course)
		}
	}

	// depend[i]是个map，表示先学i，后学map包含的课程
	depend := make([]map[int]struct{}, n)

	for q.Len() > 0 {
		elem := q.Front()
		course, _ := elem.Value.(int)
		q.Remove(elem)

		for _, preLearn := range reverseAdj[course] {
			outDegree[preLearn]--
			if outDegree[preLearn] == 0 {
				q.PushBack(preLearn)
			}
			if depend[preLearn] == nil {
				depend[preLearn] = make(map[int]struct{}, 0)
			}
			// 先学preLearn，后学course
			depend[preLearn][course] = struct{}{}
			// 注意：因为有0-1-2，学2之前也需先学0，dp[0]={1,2}
			// 需要把学course之前学习的课程加进来
			for c := range depend[course] {
				depend[preLearn][c] = struct{}{}
			}
		}
	}

	res := make([]bool, len(queries))
	for i, pair := range queries {
		if _, ok := depend[pair[0]][pair[1]]; ok {
			res[i] = true
		}
	}

	return res
}
