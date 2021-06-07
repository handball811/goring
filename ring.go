package goring

// TODO: Add Unit Test

var (
	RingDefaultSize = 8
)

type Ring struct {
	RingOp
	list []interface{}
	top  int
	bot  int
	len  int
}

func NewRing() *Ring {
	return NewRingWithSize(RingDefaultSize)
}

func NewRingWithSize(size int) *Ring {
	if size <= 0 {
		size = RingDefaultSize
	}
	return &Ring{
		list: make([]interface{}, size),
		top:  0,
		bot:  0,
		len:  0,
	}
}

func (r *Ring) Cap() int {
	return len(r.list)
}

func (r *Ring) Len() int {
	return r.len
}

func (r *Ring) Push(t interface{}) {
	if r.len+1 > r.Cap() {
		r.grow(r.Cap() * 2)
	}
	// Pushing at bottom
	r.list[r.bot] = t
	r.bot++
	r.len++
	r.bot %= r.Cap()
}

func (r *Ring) PushSlice(s []interface{}) {
	slen := len(s)
	if r.len+slen > r.Cap() {
		r.grow((r.len + slen) * 2)
	}
	// Pushing at bottom
	cap := r.Cap()
	last := r.bot + slen
	if last > cap {
		sz := cap - r.bot
		copy(r.list[r.bot:], s[:sz])
		copy(r.list, s[sz:])
	} else {
		copy(r.list[r.bot:], s)
	}
	r.bot = last % r.Cap()
	r.len += slen
}

func (r *Ring) Pop() (interface{}, bool) {
	if r.len <= 0 {
		return nil, false
	}
	// Pop
	ret := r.list[r.top]
	r.top++
	r.len--
	r.top %= r.Cap()
	return ret, true
}

func (r *Ring) PopSlice(s []interface{}) int {
	slen := len(s)
	if slen > r.len {
		slen = r.len
	}
	if slen == 0 {
		return 0
	}
	// Pop Into Slice
	cap := r.Cap()
	last := r.top + slen
	if last > cap {
		sz := cap - r.top
		copy(s[:sz], r.list[r.top:])
		copy(s[sz:], r.list)
	} else {
		copy(s, r.list[r.top:])
	}
	r.top = last % r.Cap()
	r.len -= slen
	return slen
}

func (r *Ring) At(i int) (interface{}, bool) {
	if i > r.len {
		return nil, false
	}
	return r.list[(r.top+i)%r.Cap()], true
}

func (r *Ring) Range(c func(i int, t interface{}) bool) {
	top := r.top
	for i := 0; i < r.len; i++ {
		if !c(i, r.list[top]) {
			return
		}
		top = (top + i) % r.Cap()
	}
	return
}

func (r *Ring) SetCap(c int) bool {
	return r.grow(c)
}

func (r *Ring) Clean() {
	r.top = 0
	r.bot = 0
	r.len = 0
}

func (r *Ring) CloneTo(op RingOp) {
	op.Clean()
	op.PushSlice(r.list)
}

// grows the size of slice which does not allow shrinking
func (r *Ring) grow(size int) bool {
	if r.Cap() >= size {
		return false
	}
	// generate slice
	next := make([]interface{}, size)

	// copy
	if r.bot <= r.top && r.len != 0 {
		copy(next, r.list[r.top:])
		copy(next[r.len-r.top:], r.list[:r.bot])
	} else if r.bot > r.top {
		copy(next, r.list[r.top:r.bot])
	}
	r.list = next
	r.top = 0
	r.bot = r.len
	return true
}
