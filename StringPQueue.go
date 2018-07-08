package main

type StringPQueue struct {
	data []string
	size int
}

func NewStringPQueue() *StringPQueue {
	return &StringPQueue{make([]string, 0), 0}
}

func (p *StringPQueue) less(index1, index2 int) bool {
	return p.data[index1] < p.data[index2]
}

func (p *StringPQueue) swap(index1, index2 int) {
	p.data[index1], p.data[index2] = p.data[index2], p.data[index1]
}

func father(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return i*2 + 1
}
func right(i int) int {
	return i*2 + 2
}

func (p *StringPQueue) Empty() bool {
	return p.size == 0
}

func (p *StringPQueue) Size() int {
	return p.size
}

func (p *StringPQueue) Push(s string) {
	p.data = append(p.data, s)
	pos := p.size
	f := father(pos)
	p.size++
	for pos > 0 && p.less(pos, f) {
		p.swap(pos, f)
		pos = f
		f = father(pos)
	}
}

func (p *StringPQueue) Top() string {
	return p.data[0]
}

func (p *StringPQueue) Pop() string {
	save := p.data[0]
	p.size--
	p.data[0] = p.data[p.size]
	pos := 0
	next := pos
	l := left(pos)
	r := right(pos)
	for {
		if l >= p.size {
			break
		} else {
			if p.less(l, pos) {
				p.swap(l, pos)
				next = l
			}
		}
		if r >= p.size {
			break
		} else {
			if p.less(r, pos) {
				p.swap(r, pos)
				next = r
			}
		}
		if next == pos {
			break
		} else {
			pos = next
			l = left(pos)
			r = right(pos)
		}
	}
	return save
}
