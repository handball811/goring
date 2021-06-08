package goring

import "sync"

type SyncRing struct {
	RingOp
	sync.RWMutex
	r *Ring
}

func NewSyncRing() *SyncRing {
	return &SyncRing{
		r: NewRing(),
	}
}

func (r *SyncRing) Cap() int {
	r.RLock()
	defer r.RUnlock()
	return r.r.Cap()
}

func (r *SyncRing) Len() int {
	r.RLock()
	defer r.RUnlock()
	return r.r.Len()
}

func (r *SyncRing) Push(t interface{}) {
	r.Lock()
	defer r.Unlock()
	r.r.Push(t)
}

func (r *SyncRing) PushSlice(s []interface{}) {
	r.Lock()
	defer r.Unlock()
	r.r.PushSlice(s)
}

func (r *SyncRing) Pop() (interface{}, bool) {
	r.Lock()
	defer r.Unlock()
	return r.r.Pop()
}

func (r *SyncRing) PopSlice(s []interface{}) int {
	r.Lock()
	defer r.Unlock()
	return r.r.PopSlice(s)
}

func (r *SyncRing) At(i int) (interface{}, bool) {
	r.RLock()
	defer r.RUnlock()
	return r.r.At(i)
}

func (r *SyncRing) Range(c func(i int, t interface{}) bool) {
	r.RLock()
	defer r.RUnlock()
	r.r.Range(c)
}

func (r *SyncRing) SetCap(c int) bool {
	r.Lock()
	defer r.Unlock()
	return r.r.SetCap(c)
}

func (r *SyncRing) Clean() {
	r.Lock()
	defer r.Unlock()
	r.r.Clean()
}

func (r *SyncRing) CloneTo(op RingOp) {
	r.RLock()
	defer r.RUnlock()
	r.r.CloneTo(op)
}
