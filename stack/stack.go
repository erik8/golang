package stack

const capacity = 10

// stack holding items
type Stack struct {
	size int
	data [capacity]int
}

// Pushes an item on the stack
func (s *Stack) Push(item int) {
	s.data[s.size] = item
	s.size++
}

func (s *Stack) Pop() (ret int) {
	s.size--
	ret = s.data[s.size]
	return
}
