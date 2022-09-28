package command

type Stack []Command

func NewStack() *Stack {
	stk := Stack(make([]Command, 0))
	return &stk
}

func (s *Stack) Push(cmd Command) {
	*s = append(*s, cmd)
}

func (s *Stack) Pop() Command {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Clear() {
	*s = make([]Command, 0)
}
