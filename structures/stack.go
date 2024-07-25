package structures

type Stack struct {
	Stack []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.Stack = append(s.Stack, value)
}
func (s *Stack) Pop() {
	s.Stack = s.Stack[:len(s.Stack)-1]
}
func (s *Stack) Peek() interface{} {
	return s.Stack[len(s.Stack)-1]
}
func (s *Stack) IsEmpty() bool {
	return len(s.Stack) == 0
}
func (s *Stack) Size() int {
	return len(s.Stack)
}
