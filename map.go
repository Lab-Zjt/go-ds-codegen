package main

import (
	"errors"
)

type mColor bool

const (
	mRed   mColor = true
	mBlack mColor = false
	mZero  string = ""
)

type cmpf func(float32, float32) bool

type Map struct {
	root *mapNode
	size int
	less cmpf
}

func NewMap(less func(float32, float32) bool) *Map {
	return &Map{nil, 0, less}
}

type mapNode struct {
	key    float32
	value  string
	color  mColor
	parent *mapNode
	left   *mapNode
	right  *mapNode
}

type MapIterator struct {
	n *mapNode
	m *Map
}

type MapPair struct {
	First  float32
	Second string
}

func newMapNode(k float32, v string) *mapNode {
	return &mapNode{k, v, mRed, nil, nil, nil}
}

func (n *mapNode) isLeft() bool {
	if n.parent != nil {
		if n.parent.left == n {
			return true
		}
	}
	return false
}

func (n *mapNode) isRight() bool {
	if n.parent != nil && n.parent.right == n {
		return true
	}
	return false
}

func (n *mapNode) isRed() bool {
	if n == nil {
		return false
	}
	return n.color == mRed
}

func (n *mapNode) isBlack() bool {
	if n == nil {
		return true
	}
	return n.color == mBlack
}

func (n *mapNode) mColorFlip() *mapNode {
	n.left.color = mBlack
	n.right.color = mBlack
	n.color = mRed
	return n
}

func (n *mapNode) mLeftRotate() *mapNode {
	r := n.right
	n.color, r.color = r.color, n.color
	p := n.parent
	n.right = r.left
	if r.left != nil {
		r.left.parent = n
	}
	r.left = n
	n.parent = r
	if p != nil {
		if p.left == n {
			p.left = r
		} else {
			p.right = r
		}
	}
	r.parent = p
	return r
}

func (n *mapNode) mRightRotate() *mapNode {
	l := n.left
	n.color, l.color = l.color, n.color
	p := n.parent
	n.left = l.right
	if l.right != nil {
		l.right.parent = n
	}
	l.right = n
	n.parent = l
	if p != nil {
		if p.left == n {
			p.left = l
		} else {
			p.right = l
		}
	}
	l.parent = p
	return l
}

func (n *mapNode) mTryColorFilp() *mapNode {
	r := n.right
	l := n.left
	if l.isRed() && r.isRed() {
		n = n.mColorFlip()
	}
	return n
}

func (n *mapNode) mTryLeftRotate() *mapNode {
	r := n.right
	if r.isRed() {
		n = n.mLeftRotate()
	}
	return n
}

func (n *mapNode) mTryRightRotate() *mapNode {
	l := n.left
	if l.isRed() && l.left.isRed() {
		n = n.mRightRotate()
	}
	return n
}

func (n *mapNode) mRenewal() *mapNode {
	if n == nil {
		return nil
	}
	return n.mTryColorFilp().mTryLeftRotate().mTryRightRotate().mTryColorFilp()
}

func (n *mapNode) mInsertLeft(new *mapNode) {
	n.left = new
	new.parent = n
}

func (n *mapNode) mInsertRight(new *mapNode) {
	n.right = new
	new.parent = n
}

func (n *mapNode) min() *mapNode {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *mapNode) max() *mapNode {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *mapNode) mDeleteMin() *mapNode {
	p := n.parent
	cur := n
	for {
		if cur.left == nil {
			if cur.parent != nil {
				if cur.parent.left == cur {
					cur.parent.left = nil
				} else {
					cur.parent.right = nil
				}
			}
			break
		}
		if cur.left.isBlack() && cur.left.left.isBlack() {
			cur = cur.mColorFlipOnDelete()
			if cur.right.left.isRed() {
				cur.right = cur.right.mRightRotate()
				cur = cur.mLeftRotate().mColorFlip()
			}
		}
		cur = cur.left
	}
	if cur.parent == p {
		return nil
	}
	for cur.parent != p {
		cur = cur.parent.mRenewal()
	}
	return cur.mRenewal()
}

func (n *mapNode) mDeleteMax() *mapNode {
	cur := n
	p := n.parent
	for {
		if cur.left.isRed() {
			cur = cur.mRightRotate()
		}
		if cur.right == nil {
			if cur.parent != nil {
				if cur.parent.left == cur {
					cur.parent.left = nil
				} else {
					cur.parent.right = nil
				}
			}
			break
		}
		if cur.right.isBlack() && cur.right.left.isBlack() {
			cur = cur.mColorFlipOnDelete()
			if cur.left.left.isRed() {
				cur = cur.mRightRotate().mColorFlip()
			}
		}
		cur = cur.right
	}
	if cur.parent == p {
		return nil
	}
	for cur.parent != p {
		cur = cur.parent.mRenewal()
	}
	return cur.mRenewal()
}

func (n *mapNode) mErase(k float32) (*mapNode, bool) {
	cur := n
	flagNotFound := false
	flagLast := false
	p := n.parent
	for {
		if k < cur.key {
			if cur.left == nil {
				flagNotFound = true
				break
			}
			if cur.left.isBlack() && cur.left.left.isBlack() {
				cur = cur.mColorFlipOnDelete()
				if cur.right.left.isRed() {
					cur.right = cur.right.mRightRotate()
					cur = cur.mLeftRotate().mColorFlip()
				}
			}
			cur = cur.left
		} else {
			if k > cur.key && (cur.right == nil) {
				flagNotFound = true
				break
			}
			if cur.left.isRed() {
				cur = cur.mRightRotate().right
			}
			if k == cur.key && (cur.right == nil) {
				if cur.parent != nil {
					if cur.parent.left == cur {
						cur.parent.left = nil
					} else {
						cur.parent.right = nil
					}
				} else {
					flagLast = true
				}
				break
			}
			if cur.right.isBlack() && cur.right.left.isBlack() {
				cur = cur.mColorFlipOnDelete()
				if cur.left.left.isRed() {
					cur = cur.mRightRotate().mColorFlip().right
				}
			}
			if k == cur.key {
				m := cur.right.min()
				cur.key = m.key
				cur.value = m.value
				cur.right = cur.right.mDeleteMin()
				cur = cur.mRenewal()
				break
			} else {
				cur = cur.right
			}
		}
	}
	if flagLast && !flagNotFound {
		return nil, flagNotFound
	}
	for cur.parent != p {
		cur = cur.parent.mRenewal()
	}
	return cur.mRenewal(), flagNotFound
}

func (n *mapNode) mColorFlipOnDelete() *mapNode {
	n.color = mBlack
	n.left.color = mRed
	n.right.color = mRed
	return n
}

func newMapIterator() MapIterator {
	return MapIterator{nil, nil}
}

func (it *MapIterator) Illegal() {
	it.m = nil
	it.n = nil
}

func (it MapIterator) Get() MapPair {
	return MapPair{it.n.key, it.n.value}
}

func (it MapIterator) Next() MapIterator {
	if it.n == nil {
		return it
	}
	if it.n.right != nil {
		return MapIterator{it.n.right.min(), it.m}
	}
	if it.n.isLeft() {
		return MapIterator{it.n.parent, it.m}
	}
	cur := it.n.parent
	for cur.isRight() {
		cur = cur.parent
	}
	if cur.isLeft() {
		return MapIterator{cur.parent, it.m}
	}
	return MapIterator{nil, it.m}
}

func (it MapIterator) Prev() MapIterator {
	if it.n == nil {
		return MapIterator{it.m.root.max(), it.m}
	}
	if it.n.left != nil {
		return MapIterator{it.n.left.max(), it.m}
	}
	if it.n.isRight() {
		return MapIterator{it.n.parent, it.m}
	}
	cur := it.n.parent
	for cur.isLeft() {
		cur = cur.parent
	}
	if cur.isRight() {
		return MapIterator{cur.parent, it.m}
	}
	return it
}

func (it MapIterator) Add(i int) MapIterator {
	for i > 0 {
		it = it.Next()
		i--
	}
	return it
}

func (it MapIterator) Minus(i int) MapIterator {
	for i > 0 {
		it = it.Prev()
		i--
	}
	return it
}

func (m *Map) Begin() MapIterator {
	return MapIterator{m.root.min(), m}
}

func (m *Map) End() MapIterator {
	return MapIterator{nil, m}
}

func (m *Map) mInsert(n *mapNode) {
	cur := m.root
	for {
		if m.less(n.key, cur.key) {
			if cur.left == nil {
				cur.mInsertLeft(n)
				break
			} else {
				cur = cur.left
			}
		} else if n.key == cur.key {
			cur.value = n.value
			break
		} else {
			if cur.right == nil {
				cur.mInsertRight(n)
				break
			} else {
				cur = cur.right
			}
		}
	}
	for cur.parent != nil {
		cur = cur.mRenewal().parent
	}
	m.root = cur.mRenewal()
	m.root.color = mBlack
}

func (m *Map) Insert(k float32, v string) {
	n := newMapNode(k, v)
	if m.root == nil {
		m.root = n
		m.root.color = mBlack
	} else {
		m.mInsert(n)
	}
	m.size++
}

func (m *Map) At(k float32) (string, error) {
	cur := m.root
	if cur == nil {
		return mZero, errors.New("No such key")
	}
	for {
		if k == cur.key {
			return cur.value, nil
		}
		if m.less(k, cur.key) {
			if cur.left != nil {
				cur = cur.left
				continue
			} else {
				return mZero, errors.New("No such key")
			}
		}
		if m.less(cur.key, k) {
			if cur.right != nil {
				cur = cur.right
				continue
			} else {
				return mZero, errors.New("No such key")
			}
		}
	}
}

func (m *Map) Erase(k float32) error {
	if m.size <= 0 || m.root == nil {
		return errors.New("Map is empty")
	}
	m.size--
	cur := m.root
	flagNotFound := false
	flagLast := false
	p := m.root.parent
	for {
		if m.less(k, cur.key) {
			if cur.left == nil {
				flagNotFound = true
				break
			}
			if cur.left.isBlack() && cur.left.left.isBlack() {
				cur = cur.mColorFlipOnDelete()
				if cur.right.left.isRed() {
					cur.right = cur.right.mRightRotate()
					cur = cur.mLeftRotate().mColorFlip()
				}
			}
			cur = cur.left
		} else {
			if m.less(cur.key, k) && (cur.right == nil) {
				flagNotFound = true
				break
			}
			if cur.left.isRed() {
				cur = cur.mRightRotate().right
			}
			if k == cur.key && (cur.right == nil) {
				if cur.parent != nil {
					if cur.parent.left == cur {
						cur.parent.left = nil
					} else {
						cur.parent.right = nil
					}
				} else {
					flagLast = true
				}
				break
			}
			if cur.right.isBlack() && cur.right.left.isBlack() {
				cur = cur.mColorFlipOnDelete()
				if cur.left.left.isRed() {
					cur = cur.mRightRotate().mColorFlip().right
				}
			}
			if k == cur.key {
				m := cur.right.min()
				cur.key = m.key
				cur.value = m.value
				cur.right = cur.right.mDeleteMin()
				cur = cur.mRenewal()
				break
			} else {
				cur = cur.right
			}
		}
	}
	if flagLast && !flagNotFound {
		cur = nil
		goto Result
	}
	for cur.parent != p {
		cur = cur.parent.mRenewal()
	}
	cur = cur.mRenewal()

Result:
	m.root = cur
	if m.root != nil {
		m.root.color = mBlack
	}
	if flagNotFound {
		return errors.New("No such key")
	}
	return nil
}

func (m *Map) Size() int {
	return m.size
}

type _MapFunc func(v string) string

func (m *Map) foreach(n *mapNode, f _MapFunc) {
	if n == nil {
		return
	}
	m.foreach(n.left, f)
	n.value = f(n.value)
	m.foreach(n.right, f)
}

type _FilterFunc func(v string) bool

func (m *Map) filter(n *mapNode, nm *Map, f _FilterFunc) {
	if n == nil {
		return
	}
	m.filter(n.left, nm, f)
	if f(n.value) {
		nm.Insert(n.key, n.value)
	}
	m.filter(n.right, nm, f)
}

func (m *Map) Filter(f _FilterFunc) *Map {
	nm := NewMap(m.less)
	m.filter(m.root, nm, f)
	return nm
}
func (m *Map) ForEach(f _MapFunc) {
	if m.root == nil {
		return
	}
	m.foreach(m.root, f)
}

func (m *Map) mapping(n *mapNode, nm *Map, f _MapFunc) {
	if n == nil {
		return
	}
	m.mapping(n.left, nm, f)
	nm.Insert(n.key, f(n.value))
	m.mapping(n.right, nm, f)
}

func (m *Map) Map(f _MapFunc) *Map {
	nm := NewMap(m.less)
	m.mapping(m.root, nm, f)
	return nm
}

type _ReduceFunc func(node1, node2 string) string

func (m *Map) reduce(n *mapNode, res string, f _ReduceFunc) string {
	if n == nil {
		return res
	}
	res = m.reduce(n, res, f)
	res = f(res, n.value)
	res = m.reduce(n, res, f)
	return res
}

func (m *Map) Reduce(f _ReduceFunc) string {
	var res string
	return m.reduce(m.root, res, f)
}
