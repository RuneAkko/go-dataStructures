package stack

// Stack 利用数组实现基本的LIFO栈
type Stack struct {
	size int //栈容量，可以容纳的栈元素个数
	S    []int
	top  int
}

// NewStack 返回一个新的栈
func NewStack(length int) *Stack {
	s := &Stack{size: length}
	s.S = make([]int, length+1)
	s.top = 0
	return s
}

func (s *Stack) isFull() bool {
	return s.top == s.size
}

func (s *Stack) isEmpty() bool {
	return s.top == 0
}

func (s *Stack) push(ele int) {
	if s.isFull() {
		panic("stack is full")
	}
	s.top++
	s.S[s.top] = ele
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		panic("stack is empty")
	}
	ele := s.S[s.top]
	s.top--
	return ele
}
