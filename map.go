package main

import "fmt"

type mColor bool
type mDirect bool

const (
	mRed   mColor  = true
	mBlack mColor  = false
	mLeft  mDirect = true
	mRight mDirect = false
)

type mapNode struct {
	key    float32
	value  string
	color  mColor
	parent *mapNode
	left   *mapNode
	right  *mapNode
}

func newMapNode(k float32, v string) *mapNode {
	return &mapNode{k, v, mRed, nil, nil, nil}
}

type Map struct {
	root *mapNode
	size int
}

func (m *Map) less(node1, node2 *mapNode) bool {
	return node1.key < node2.key
}

func NewMap() *Map {
	return &Map{nil, 0}
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

func (m *Map) mInsert(n *mapNode) {
	cur := m.root
	for {
		if m.less(n, cur) {
			if cur.left == nil {
				cur.mInsertLeft(n)
				break
			} else {
				cur = cur.left
			}
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
	/*if n.left == nil {
		return nil
	}
	if n.left.isBlack() && n.left.left.isBlack() {
		n = n.mColorFlip()
		if n.right.left.isRed() {
			n = n.mLeftRotate().mColorFlip()
		}
	}
	n.left = n.left.mDeleteMin()*/
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
	return cur
}

func (n *mapNode) mDeleteMax() *mapNode {
	/*if n.left.isRed() {
		n = n.mRightRotate()
	}
	if n.right == nil {
		return nil
	}
	if n.right.isBlack() && n.right.left.isBlack() {
		n = n.mColorFlipOnDelete()
		if n.left.left.isBlack() {
			n = n.mRightRotate().mColorFlip()
		}
	}
	n.left = n.left.mDeleteMin()*/

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
	return cur
}

func (n *mapNode) mErase(k float32) *mapNode {
	/*if k < n.key {
		if n.left.isBlack() && n.left.left.isBlack() {
			n = n.mColorFlip()
			if n.right.left.isRed() {
				n.right = n.right.mRightRotate()
				n = n.mLeftRotate().mColorFlip()
			}
		}
		n.left = n.left.mErase(k)
	} else {
		if n.left.isRed() {
			n = n.mRightRotate()
		}
		if k == n.key && (n.right == nil) {
			return nil
		}
		if n.right.isBlack() && n.right.left.isBlack() {
			n = n.mColorFlipOnDelete()
			if n.left.left.isBlack() {
				n = n.mRightRotate().mColorFlip()
			}
		}
		if k == n.key {
			m := n.right.min()
			n.value = m.value
			n.key = m.key
			n.right = n.right.mDeleteMin()
		} else {
			n.right = n.right.mErase(k)
		}
	}
	return n.mRenewal()*/
	cur := n
	p := n.parent
	for {
		if k < cur.key {
			if cur.left.isBlack() && cur.left.left.isBlack() {
				cur = cur.mColorFlip()
				if cur.right.left.isRed() {
					cur.right = cur.right.mRightRotate()
					cur = cur.mLeftRotate().mColorFlip()
				}
			}
			cur = cur.left
		} else {
			if cur.left.isRed() {
				cur = cur.mRightRotate()
			}
			if k == cur.key && (cur.right == nil) {
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
				cur = cur.mColorFlip()
				if cur.left.left.isBlack() {
					cur = cur.mRightRotate().mColorFlip()
				}
			}
			if k == cur.key {
				m := cur.right.min()
				cur.key = m.key
				cur.value = m.value
				cur.right = cur.right.mDeleteMin()
			} else {
				cur = cur.right
			}
		}
	}
	if cur.parent == p {
		return nil
	}
	for cur.parent != p {
		cur = cur.parent.mRenewal()
	}
	return cur
}

func (n *mapNode) mColorFlipOnDelete() *mapNode {
	n.color = mBlack
	n.left.color = mRed
	n.right.color = mRed
	return n
}

func (m *Map) deleteMin() {
	m.root = m.root.mDeleteMin()
	if m.root == nil {
		return
	}
	m.root.color = mBlack
}

func (m *Map) deleteMax() {
	m.root = m.root.mDeleteMax()
	if m.root == nil {
		return
	}
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

func (m *Map) Size() int {
	return m.size
}

type MapFunc func(k float32, v string)

func (m *Map) foreach(n *mapNode, f MapFunc) {
	if n == nil {
		return
	}
	m.foreach(n.left, f)
	f(n.key, n.value)
	m.foreach(n.right, f)
}

func (m *Map) ForEach(f MapFunc) {
	if m.root == nil {
		return
	}
	m.foreach(m.root, f)
}

func (m *Map) At(k float32) string {
	cur := m.root
	if cur == nil {
		return ""
	}
	for {
		if k == cur.key {
			return cur.value
		}
		if k < cur.key {
			if cur.left != nil {
				cur = cur.left
				continue
			} else {
				return ""
			}
		}
		if k > cur.key {
			if cur.right != nil {
				cur = cur.right
				continue
			} else {
				return ""
			}
		}
	}
}

func (m *Map) Erase(k float32) {

}

func PrintMap(m *Map) {
	if m == nil {
		return
	}
	printMap(m.root)
}

func printMap(cur *mapNode) {
	if cur == nil {
		return
	}
	printMap(cur.left)
	var lkey, rkey, pkey float32
	if cur.left != nil {
		lkey = cur.left.key
	}
	if cur.right != nil {
		rkey = cur.right.key
	}
	if cur.parent != nil {
		pkey = cur.parent.key
	}
	var c string
	if cur.color == mRed {
		c = "red"
	} else {
		c = "black"
	}
	fmt.Printf("key:%.2f color:%s left:%.2f right:%.2f parent:%.2f\n", cur.key, c, lkey, rkey, pkey)
	printMap(cur.right)
}
