package main

import (
	"errors"
)

type node struct {
	data       string
	next, prev *node
}

type StringList struct {
	size       int
	head, tail *node
}

type ListIt struct {
	n *node
	l *StringList
}

func NewStringList() *StringList {
	return &StringList{size: 0, head: nil, tail: nil}
}

func (l *StringList) Empty() bool {
	return l.size > 0
}

func (l *StringList) Size() int {
	return l.size
}

func (l *StringList) PushBack(d string) {
	if l.head == nil {
		l.head = &node{d, nil, nil}
		l.tail = l.head
		l.size++
	} else {
		n := &node{d, nil, l.tail}
		l.tail.next = n
		l.tail = n
		l.size++
	}
}

func (l *StringList) PushFront(d string) {
	if l.tail == nil {
		l.head = &node{d, nil, nil}
		l.tail = l.head
		l.size++
	} else {
		n := &node{d, l.head, nil}
		l.head.prev = n
		l.head = n
		l.size++
	}
}
func (l *StringList) PopBack() {
	if l.tail == nil {
		return
	}
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		l.size = 0
		return
	}
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	}
	l.size--
}

func (l *StringList) PopFront() {
	if l.head == nil {
		return
	}
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		l.size = 0
		return
	}
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}
	l.size--
}

func (i ListIt) Next() ListIt {
	return ListIt{i.n.next, i.l}
}

func (i ListIt) Prev() ListIt {
	if i.n == nil {
		return ListIt{i.l.tail, i.l}
	}
	return ListIt{i.n.prev, i.l}
}

func (i ListIt) Add(d int) ListIt {
	if d <= 0 || i.n == nil {
		return i
	}
	for d > 0 {
		if i.n.next == nil {
			return i
		}
		d--
		i.n = i.n.next
	}
	return i
}

func (i ListIt) Minus(d int) ListIt {
	if d <= 0 {
		return i
	}
	if i.n == nil {
		d--
		i.n = i.l.tail
	}
	for d > 0 {
		if i.n.prev == nil {
			return i
		}
		d--
		i.n = i.n.prev
	}
	return i
}

func (i ListIt) Get() string {
	return i.n.data
}

func (l *StringList) Begin() ListIt {
	return ListIt{l.head, l}
}

func (l *StringList) End() ListIt {
	return ListIt{nil, l}
}
func (l *StringList) Front() string {
	return l.head.data
}

func (l *StringList) Back() string {
	return l.tail.data
}
func (l *StringList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}
func (l *StringList) Insert(it ListIt, d string) error {
	if l != it.l {
		return errors.New("Imcompitible StringList")
	}
	if it.n == nil {
		l.PushBack(d)
	} else {
		n := &node{d, it.n, nil}
		if it.n.prev != nil {
			n.prev = it.n.prev
			it.n.prev.next = n
		} else {
			l.head = n
		}
		it.n.prev = n
		l.size++
	}
	return nil
}
func (l *StringList) Erase(it ListIt) error {
	if l != it.l {
		return errors.New("Imcompitible StringList")
	}
	if it.n == nil {
		return nil
	}
	next := it.n.next
	prev := it.n.prev
	if next != nil {
		it.n.next.prev = prev
	} else {
		l.tail = prev
	}
	if prev != nil {
		it.n.prev.next = next
	} else {
		l.head = next
	}
	it.n = nil
	l.size--
	return nil
}

func (l *StringList) Find(d string) ListIt {
	for it := l.Begin(); it != l.End(); it = it.Next() {
		if it.Get() == d {
			return it
		}
	}
	return l.End()
}
