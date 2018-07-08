package main

type Queue struct {
	data []string
	size int
	head int
	tail int
}

func NewQueue() *Queue {
	return &Queue{make([]string, 0), 0, 0, 0}
}

func (q *Queue) Empty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Push(d string) {
	if q.head >= q.size {
		q.data = q.data[q.head:q.tail]
		q.head = 0
		q.tail = q.size
	}
	q.data = append(q.data, d)
	q.tail++
	q.size++
}

func (q *Queue) Pop() string {
	q.head++
	q.size--
	return q.data[q.head-1]
}

func (q *Queue) Front() string {
	return q.data[q.head]
}

func (q *Queue) Back() string {
	return q.data[q.tail-1]
}

