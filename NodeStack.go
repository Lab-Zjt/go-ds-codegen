package main

type nodeStack struct {
	data []*mapNode
	size int
	ptr  int
}

func NewnodeStack() *nodeStack {
	return &nodeStack{make([]*mapNode, 0), 0, 0}
}

func (s *nodeStack) Empty() bool {
	return s.size == 0
}

func (s *nodeStack) Size() int {
	return s.size
}

func (s *nodeStack) Top() *mapNode {
	return s.data[s.ptr-1]
}

func (s *nodeStack) Push(d *mapNode) {
	s.data = append(s.data, d)
	s.ptr++
	s.size++
}

func (s *nodeStack) Pop() *mapNode {
	s.ptr--
	s.size--
	return s.data[s.ptr]
}

func (s *nodeStack) Clear() {
	s.ptr = 0
	s.size = 0
	s.data = make([]*mapNode, 0)
}

