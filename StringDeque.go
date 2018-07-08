package main

type StringDeque struct {
	data []string
	size int
	head int
	tail int
}

func NewStringDeque() *StringDeque {
	return &StringDeque{make([]string, 0), 0, 0, 0}
}

func (d *StringDeque) Empty() bool {
	return d.size == 0
}

func (d *StringDeque) Size() int {
	return d.size
}

func (d *StringDeque) PushBack(s string) {
	if d.head > d.size {
		d.data = d.data[d.head:d.tail]
		d.head = 0
		d.tail = d.size
	}
	d.data = append(d.data, s)
	d.tail++
	d.size++
}

func (d *StringDeque) PushFront(s string) {
	//TODO: Optimize resize
	if d.head == 0 {
		d.data = append(make([]string, d.size+1), d.data...)
		d.head += d.size + 1
		d.tail += d.size + 1
	}
	d.head--
	d.data[d.head] = s
	d.size++
}

func (d *StringDeque) PopBack() string {
	d.tail--
	d.size--
	return d.data[d.tail]
}

func (d *StringDeque) PopFront() string {
	d.head++
	d.size--
	return d.data[d.head-1]
}

func (d *StringDeque) Clear() {
	d.data = make([]string, 0)
	d.size = 0
	d.head = 0
	d.tail = 0
}

