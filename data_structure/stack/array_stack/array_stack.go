package array_stack

type arrayStack struct {
	list []string
	size int
}

func InitArrayStack() *arrayStack {
	stack := &arrayStack{
		list: make([]string, 0),
		size:0,
	}
	return stack
}

func (s *arrayStack) Push(data string) {
	s.list = append(s.list, data)
	s.size++
}

func (s *arrayStack) Pop() (string, bool) {
	if s.size <= 0 {
		return "", false
	}
	node := s.list[s.size-1]
	s.list[s.size-1] = ""
	s.size--
	return node, true
}

func (s *arrayStack) Empty() bool {
	return s.size <= 0
}