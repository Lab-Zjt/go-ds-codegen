package main

type StringVector struct {
	data []string
}

type VectorIt struct {
	off int
	v   *StringVector
}

func (i VectorIt) Get() string {
	return i.v.data[i.off]
}
func (i VectorIt) Next() VectorIt {
	return VectorIt{i.off + 1, i.v}
}

func (i VectorIt) Prev() VectorIt {
	return VectorIt{i.off - 1, i.v}
}

func (i VectorIt) Add(d int) VectorIt {
	if i.off+d > len(i.v.data) {
		return VectorIt{len(i.v.data), i.v}
	}
	return VectorIt{i.off + d, i.v}
}

func (i VectorIt) Minus(d int) VectorIt {
	if i.off-d < 0 {
		return VectorIt{0, i.v}
	}
	return VectorIt{i.off - d, i.v}
}

func NewStringVector() *StringVector {
	return &StringVector{make([]string, 0)}
}

func (v *StringVector) Empty() bool {
	return len(v.data) == 0
}

func (v *StringVector) Size() int {
	return len(v.data)
}

func (v *StringVector) Begin() VectorIt {
	return VectorIt{0, v}
}

func (v *StringVector) End() VectorIt {
	return VectorIt{len(v.data), v}
}

func (v *StringVector) PushBack(s string) {
	v.data = append(v.data, s)
}

func (v *StringVector) Data() []string {
	return v.data
}

func (v *StringVector) PopBack() {
	v.data = v.data[0 : len(v.data)-1]
}

func (v *StringVector) Front() string {
	return v.data[0]
}

func (v *StringVector) Back() string {
	return v.data[len(v.data)-1]
}

func (v *StringVector) At(i int) string {
	return v.data[i]
}
