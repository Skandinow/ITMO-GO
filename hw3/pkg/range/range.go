package rangeI

import "strconv"

type RangeInt interface {
	Length() int
	Intersect(other RangeInt)
	Union(other RangeInt) bool
	IsEmpty() bool
	ContainsInt(i int) bool
	ContainsRange(other RangeInt) bool
	IsIntersect(other RangeInt) bool
	ToSlice() []int
	Minimum() (int, bool)
	Maximum() (int, bool)
	String() string
}

type Range struct {
	start int
	end   int
}

func NewRangeInt(a, b int) *Range {
	return &Range{start: a, end: b}
}

func (r *Range) Length() int {
	if r.IsEmpty() {
		return 0
	}
	return r.end - r.start + 1
}

func (r *Range) Minimum() (int, bool) {
	if r.IsEmpty() {
		return 0, false
	}

	return r.start, true

}

func (r *Range) Maximum() (int, bool) {
	if r.IsEmpty() {
		return 0, false
	} else {
		return r.end, true
	}
}

func (r *Range) IsEmpty() bool {
	return r.end < r.start
}

func (r *Range) Intersect(other RangeInt) {
	minOther, _ := other.Minimum()
	maxOther, _ := other.Maximum()

	if r.IsEmpty() || minOther > maxOther || r.start > maxOther || minOther > r.end {
		*r = *NewEmptyRange()
	} else if r.start <= minOther && r.end >= maxOther {
		r.start = minOther
		r.end = maxOther
	} else if r.start >= minOther && r.start <= maxOther {
		r.end = maxOther
	} else if minOther >= r.start && minOther <= r.end {
		r.start = minOther
	}
}

func (r *Range) Union(other RangeInt) bool {
	minOther, _ := other.Minimum()
	maxOther, _ := other.Maximum()

	if r.IsEmpty() && other.IsEmpty() || other.IsEmpty() && r.start <= r.end {
		return true
	} else if !other.IsEmpty() && r.start > r.end {
		r.start = minOther
		r.end = maxOther
	} else if r.IsEmpty() || r.start > maxOther+1 || minOther > r.end+1 {
		return false
	} else if r.start >= minOther && r.end <= maxOther {
		r.start = minOther
		r.end = maxOther
	} else if r.start >= minOther && r.start-1 <= maxOther {
		r.start = minOther
	} else if minOther >= r.start && minOther-1 <= r.end {
		r.end = maxOther
	}

	return true
}

func (r *Range) ContainsInt(i int) bool {
	return !r.IsEmpty() && i >= r.start && i <= r.end
}

func (r *Range) ContainsRange(other RangeInt) bool {
	if r.IsEmpty() || other.IsEmpty() {
		return other.IsEmpty()
	}

	minOther, _ := other.Minimum()
	maxOther, _ := other.Maximum()
	return minOther >= r.start && maxOther <= r.end
}

func (r *Range) IsIntersect(other RangeInt) bool {
	if r.IsEmpty() {
		return false
	}
	minOther, _ := other.Minimum()
	maxOther, _ := other.Maximum()
	return !(r.start > maxOther || r.end < minOther)
}

func (r *Range) ToSlice() []int {
	if r.IsEmpty() {
		return []int{}
	}
	result := make([]int, r.end-r.start+1)
	cnt := 0

	for i := r.start; i <= r.end; i++ {
		result[cnt] = i
		cnt++
	}
	return result
}

func (r *Range) String() string {
	if r.IsEmpty() {
		return ""
	}
	result := "[" + strconv.Itoa(r.start) + "," + strconv.Itoa(r.end) + "]"
	return result
}

func NewEmptyRange() *Range {
	return NewRangeInt(1, 0)
}
