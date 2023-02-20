package mystack

type Stack struct {
	slice []any
	size  int
}

func NewStack() (stack *Stack) {
	stack = &Stack{
		slice: make([]any, 1),
	}
	return
}
func (s *Stack) Push(data any) {
	if s.size >= len(s.slice) {
		temp := make([]any, len(s.slice)*2)
		copy(temp, s.slice)
		s.slice = temp
	}
	s.slice[s.size] = data
	s.size++
}
func (s *Stack) IsEmpty() (bool bool) {
	if s.size <= 0 {
		bool = true
	} else {
		bool = false
	}
	return
}
func (s *Stack) Pop() (data any) {
	if !s.IsEmpty() {
		s.size--
		data = s.slice[s.size]
	} else {
		data = nil
	}
	return

}
func (s *Stack) Top() (data any) {
	data = s.slice[s.size-1]
	return
}
