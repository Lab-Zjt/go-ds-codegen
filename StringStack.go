package main

type StringStack struct {
	data []string
	size int
	ptr  int
}

func NewStringStack() *StringStack {
	return &StringStack{make([]string, 0), 0, 0}
}

func (s *StringStack) Empty() bool {
	return s.size == 0
}

func (s *StringStack) Size() int {
	return s.size
}

func (s *StringStack) Top() string {
	return s.data[s.ptr-1]
}

func (s *StringStack) Push(d string) {
	s.data = append(s.data, d)
	s.ptr++
	s.size++
}

func (s *StringStack) Pop() string {
	s.ptr--
	s.size--
	return s.data[s.ptr]
}

func (s *StringStack) Clear() {
	s.ptr = 0
	s.size = 0
	s.data = make([]string, 0)
}

